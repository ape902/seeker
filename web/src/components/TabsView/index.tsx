import React, { useState, useEffect, useCallback, useRef } from 'react';
import { Tabs, Badge, message } from 'antd';
import { useNavigate, useLocation } from 'react-router-dom';
import { CloseOutlined, ReloadOutlined } from '@ant-design/icons';
import { MenuItem } from '../../config/menuConfig';
import './style.css';

export interface TabItem {
  key: string;
  label: string;
  closable: boolean;
}

interface TabsViewProps {
  menuItems: MenuItem[];
}

const findMenuItemByKey = (items: MenuItem[], key: string): MenuItem | null => {
  for (const item of items) {
    if (item.children) {
      for (const child of item.children) {
        if (child.key === key) {
          return child;
        }
      }
    }
  }
  return null;
};

const TabsView: React.FC<TabsViewProps> = ({ menuItems }) => {
  const [activeKey, setActiveKey] = useState<string>('');
  const [tabs, setTabs] = useState<TabItem[]>([]);
  const navigate = useNavigate();
  const location = useLocation();

  // 从localStorage加载标签状态
  useEffect(() => {
    const savedTabs = localStorage.getItem('tabs');
    const currentPath = location.pathname;
    
    // 始终优先使用验证后的当前路径
    const validActiveKey = menuItems.some(item => 
      item.children?.some(child => child.key === currentPath)
    ) ? currentPath : '/';

    // 同步更新activeKey和localStorage
    setActiveKey(validActiveKey);
    localStorage.setItem('activeTabKey', validActiveKey);

    if (savedTabs) {
      setTabs(JSON.parse(savedTabs));
    }
  }, [location.pathname, menuItems]); // 添加依赖项

  // 当路由变化时更新标签
  useEffect(() => {
    const currentPath = location.pathname;
    
    // 查找当前路径对应的菜单项
    const menuItem = findMenuItemByKey(menuItems, currentPath);
    
    if (menuItem) {
      // 检查标签是否已存在
      const isExist = tabs.some(tab => tab.key === currentPath);
      
      if (!isExist) {
        // 添加新标签
        const newTab: TabItem = {
          key: currentPath,
          label: menuItem.label,
          closable: tabs.length > 0
        };
        
        const newTabs = [...tabs, newTab];
        setTabs(newTabs);
        localStorage.setItem('tabs', JSON.stringify(newTabs));
      }
      
      // 更新当前活动标签
      setActiveKey(currentPath);
      localStorage.setItem('activeTabKey', currentPath);
    }
  }, [location.pathname, menuItems]);

  // 处理标签切换
  const handleTabChange = (key: string) => {
    // 先更新路由再设置状态
    navigate(key, { replace: true });
    
    // 同步更新菜单选中状态
    setActiveKey(key);
    localStorage.setItem('activeTabKey', key);
    localStorage.setItem('selectedKeys', JSON.stringify([key]));
  };

  // 处理标签关闭
  const handleTabClose = useCallback((targetKey: string) => {
    // 找到要关闭的标签的索引
    const targetIndex = tabs.findIndex(tab => tab.key === targetKey);
    
    // 创建新的标签数组
    const newTabs = tabs.filter(tab => tab.key !== targetKey);
    
    // 更新标签列表
    setTabs(newTabs);
    localStorage.setItem('tabs', JSON.stringify(newTabs));
    
    // 如果关闭的是当前活动标签，则切换到其他标签
    if (targetKey === activeKey && newTabs.length) {
      // 优先选择右侧标签，如果没有则选择左侧标签
      const newActiveKey = newTabs[targetIndex]?.key || newTabs[targetIndex - 1]?.key;
      setActiveKey(newActiveKey);
      localStorage.setItem('activeTabKey', newActiveKey);
      navigate(newActiveKey, { replace: true });
    }
  }, [activeKey, navigate, tabs]);

  // 处理标签刷新
  const handleTabRefresh = (key: string) => {
    message.success('刷新页面: ' + key);
    // 实际应用中可以在这里添加页面刷新逻辑
    navigate(key, { replace: true });
  };

  // 自定义标签项
  const renderTabItem = (tab: TabItem) => ({
    label: (
      <div className="tab-item">
        <span>{tab.label}</span>
        <div className="tab-actions">
          <ReloadOutlined
            className="tab-action-icon tab-refresh-icon"
            onClick={(e) => {
              e.stopPropagation();
              handleTabRefresh(tab.key);
            }}
          />
          {tab.closable && (
            <CloseOutlined
              className="tab-action-icon tab-close-icon"
              onClick={(e) => {
                e.stopPropagation();
                handleTabClose(tab.key);
              }}
            />
          )}
        </div>
      </div>
    ),
    key: tab.key,
    closable: tab.closable,
  });

  return (
    <div className="tabs-view-container">
      {tabs.length > 0 && (
        <Tabs
          activeKey={activeKey}
          onChange={handleTabChange}
          type="card"
          items={tabs.map(renderTabItem)}
          className="custom-tabs"
        />
      )}
    </div>
  );
};

export default TabsView;