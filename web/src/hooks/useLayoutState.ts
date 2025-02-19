import { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

export interface UseLayoutStateReturn {
  collapsed: boolean;
  setCollapsed: (value: boolean) => void;
  selectedKeys: string[];
  setSelectedKeys: (keys: string[]) => void;
  openKeys: string[];
  setOpenKeys: (keys: string[]) => void;
  userName: string;
  handleMenuClick: (params: { key: string }) => void;
  handleOpenChange: (keys: string[]) => void;
  handleLogout: () => void;
}

export const useLayoutState = (): UseLayoutStateReturn => {
  const navigate = useNavigate();
  const location = useLocation();
  const [collapsed, setCollapsed] = useState(false);
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

    let newSelectedKeys: string[];
    if (savedSelectedKeys) {
      newSelectedKeys = JSON.parse(savedSelectedKeys);
      setSelectedKeys(newSelectedKeys);
      if (currentPath !== newSelectedKeys[0]) {
        navigate(newSelectedKeys[0], { replace: true });
      }
    } else {
      newSelectedKeys = [currentPath];
      setSelectedKeys(newSelectedKeys);
      localStorage.setItem('selectedKeys', JSON.stringify(newSelectedKeys));
    }

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

  return {
    collapsed,
    setCollapsed,
    selectedKeys,
    setSelectedKeys,
    openKeys,
    setOpenKeys,
    userName,
    handleMenuClick,
    handleOpenChange,
    handleLogout
  };
};