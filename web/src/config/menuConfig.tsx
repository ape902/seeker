import React from 'react';
import {
  UserOutlined,
  LaptopOutlined,
  CloudServerOutlined,
} from '@ant-design/icons';

export interface MenuItem {
  key: string;
  icon?: React.ReactNode;
  label: string;
  children?: MenuItem[];
}

export const menuItems: MenuItem[] = [{
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