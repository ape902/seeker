import React, { useState, useEffect } from 'react';
import { Layout, Menu, ConfigProvider, theme, Avatar, Dropdown } from 'antd';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import {
  UserOutlined,
  LaptopOutlined,
  CloudServerOutlined,
  LogoutOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
} from '@ant-design/icons';
import { useLayoutState } from '../hooks/useLayoutState';
import { menuItems } from '../config/menuConfig';
import TabsView from '../components/TabsView';
import '../styles/layout.css';

const { Header, Content, Sider } = Layout;

const BasicLayout: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [collapsed, setCollapsed] = useState(false);
  const { token } = theme.useToken();
  const [selectedKeys, setSelectedKeys] = useState<string[]>([]);
  const [openKeys, setOpenKeys] = useState<string[]>([]);
  const [userName, setUserName] = useState<string>('');

  useEffect(() => {
    const savedSelectedKeys = localStorage.getItem('selectedKeys');
    const savedOpenKeys = localStorage.getItem('openKeys');
    const currentPath = location.pathname;
    const userInfo = localStorage.getItem('userInfo');

    try {
      if (userInfo) {
        const parsedUserInfo = JSON.parse(userInfo);
        setUserName(parsedUserInfo?.nick_name || '未知用户');
      } else {
        setUserName('未知用户');
      }
    } catch (error) {
      console.error('解析用户信息失败:', error);
      setUserName('未知用户');
    }

    // 处理菜单选中状态
    let newSelectedKeys: string[];
    if (savedSelectedKeys) {
      newSelectedKeys = JSON.parse(savedSelectedKeys);
      setSelectedKeys(newSelectedKeys);
      // 如果当前路径与保存的路径不一致，则导航到保存的路径
      if (currentPath !== newSelectedKeys[0]) {
        navigate(newSelectedKeys[0], { replace: true });
      }
    } else {
      newSelectedKeys = [currentPath];
      setSelectedKeys(newSelectedKeys);
      localStorage.setItem('selectedKeys', JSON.stringify(newSelectedKeys));
    }

    // 处理菜单展开状态
    if (savedOpenKeys) {
      setOpenKeys(JSON.parse(savedOpenKeys));
    } else {
      const defaultOpenKey = currentPath.startsWith('/users') ? 'user' :
                            currentPath.startsWith('/host') || 
                            currentPath.startsWith('/sftp') || 
                            currentPath.startsWith('/command') || 
                            currentPath.startsWith('/discovery') ? 'cmdb' :
                            currentPath.startsWith('/minio') ? 'storage' : 'user';
      setOpenKeys([defaultOpenKey]);
      localStorage.setItem('openKeys', JSON.stringify([defaultOpenKey]));
    }
  }, [location.pathname, navigate]);

  const menuItems = [{
    key: 'user',
    icon: <UserOutlined />,
    label: '用户中心',
    children: [
      { key: '/users', label: '用户列表' },
    ],
  },
  {
    key: 'cmdb',
    icon: <LaptopOutlined />,
    label: 'CMDB管理',
    children: [
      { key: '/hosts', label: '主机管理' },
      { key: '/sftp', label: '文件管理' },
      { key: '/command', label: '命令执行' },
      { key: '/discovery', label: '服务发现' },
    ],
  },
  {
    key: 'storage',
    icon: <CloudServerOutlined />,
    label: '存储管理',
    children: [
      { key: '/minio', label: 'MinIO管理' },
    ],
  }];

  const handleMenuClick = ({ key }: { key: string }) => {
    setSelectedKeys([key]);
    localStorage.setItem('selectedKeys', JSON.stringify([key]));
    navigate(key);
  };

  const handleOpenChange = (keys: string[]) => {
    setOpenKeys(keys);
    localStorage.setItem('openKeys', JSON.stringify(keys));
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('selectedKeys');
    localStorage.removeItem('openKeys');
    localStorage.removeItem('userInfo');
    navigate('/login');
  };

  return (
    <ConfigProvider
      theme={{
        algorithm: theme.defaultAlgorithm,
        token: {
          colorPrimary: 'var(--primary-color)',
          borderRadius: 'var(--border-radius)',
          colorBgContainer: '#ffffff',
          colorBgLayout: '#f5f7fa',
          boxShadow: 'var(--box-shadow)',
          transition: 'var(--transition-duration)',
        },
      }}
    >
      <Layout style={{ minHeight: '100vh' }}>
        <Header style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          padding: '0 32px',
          background: token.colorBgContainer,
          boxShadow: 'var(--box-shadow)',
          height: '64px',
          position: 'sticky',
          top: 0,
          zIndex: 1000,
          transition: 'all var(--transition-duration)',
        }}>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            fontSize: '20px',
            fontWeight: 600,
            color: token.colorTextHeading,
            letterSpacing: '0.5px',
          }}>
            <div
              onClick={() => setCollapsed(!collapsed)}
              style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                width: '40px',
                height: '40px',
                cursor: 'pointer',
                borderRadius: '8px',
                transition: 'all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1)',
                backgroundColor: collapsed ? 'rgba(24, 144, 255, 0.1)' : 'transparent',
                '&:hover': {
                  backgroundColor: 'rgba(24, 144, 255, 0.15)',
                  transform: 'scale(1.05)'
                }
              }}
            >
              {collapsed ? (
                <MenuUnfoldOutlined style={{ fontSize: '18px', color: token.colorPrimary }} />
              ) : (
                <MenuFoldOutlined style={{ fontSize: '18px', color: token.colorPrimary }} />
              )}
            </div>
            <span style={{ marginLeft: '12px' }}>Seeker管理系统</span>
          </div>
          <Dropdown menu={{
            items: [{
              key: 'logout',
              icon: <LogoutOutlined />,
              label: '退出登录',
              onClick: handleLogout,
              style: {
                padding: '8px 16px',
                transition: 'all var(--transition-duration)',
              }
            }],
            style: {
              padding: '4px',
              borderRadius: 'var(--border-radius)',
              boxShadow: 'var(--box-shadow)',
            }
          }}>
            <div style={{ display: 'flex', alignItems: 'center', cursor: 'pointer' }}>
              <Avatar icon={<UserOutlined />} />
              <span style={{ marginLeft: '8px' }}>{userName}</span>
            </div>
          </Dropdown>
        </Header>
        <Layout>
          <Sider
            width={220}
            collapsed={collapsed}
            style={{
              background: token.colorBgContainer,
              boxShadow: 'var(--box-shadow)',
              height: 'calc(100vh - 64px)',
              position: 'fixed',
              left: 0,
              top: 64,
              zIndex: 100,
              transition: 'all var(--transition-duration)',
            }}
          >
            <Menu
              mode="inline"
              selectedKeys={selectedKeys}
              openKeys={openKeys}
              onOpenChange={handleOpenChange}
              onClick={handleMenuClick}
              items={menuItems}
              style={{
                height: '100%',
                borderRight: 0,
              }}
            />
          </Sider>
          <Layout style={{
            marginLeft: collapsed ? 80 : 220,
            transition: 'margin-left var(--transition-duration)',
            padding: '24px',
          }}>
            {/* 添加标签卡视图 */}
            <TabsView menuItems={menuItems} />
            <Content style={{
              background: token.colorBgContainer,
              borderRadius: 'var(--border-radius)',
              padding: '24px',
              minHeight: 280,
              boxShadow: 'var(--box-shadow)',
            }}>
              <Outlet />
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </ConfigProvider>
  );
};

export default BasicLayout;