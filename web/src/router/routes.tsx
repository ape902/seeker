import React from 'react';
import { Navigate, RouteObject } from 'react-router-dom';
import BasicLayout from '@layouts/BasicLayout';

// 懒加载页面组件
const Login = React.lazy(() => import('@pages/Login'));
const Users = React.lazy(() => import('@pages/Users'));
const Hosts = React.lazy(() => import('@pages/Hosts'));
const Sftp = React.lazy(() => import('@pages/Sftp'));
const Command = React.lazy(() => import('@pages/Command'));
const Discovery = React.lazy(() => import('@pages/Discovery'));
const Minio = React.lazy(() => import('@pages/Minio'));

// 路由守卫组件
const AuthGuard: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const token = localStorage.getItem('token');
  if (!token) {
    return <Navigate to="/login" replace />;
  }
  return <>{children}</>;
};

// 路由配置
const routes: RouteObject[] = [
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/',
    element: (
      <AuthGuard>
        <BasicLayout />
      </AuthGuard>
    ),
    children: [
      {
        path: '',
        element: (() => {
          const savedSelectedKeys = localStorage.getItem('selectedKeys');
          const defaultPath = savedSelectedKeys ? JSON.parse(savedSelectedKeys)[0] : '/users';
          return <Navigate to={defaultPath} replace />;
        })(),
      },
      {
        path: 'users',
        element: <Users />,
      },
      {
        path: 'hosts',
        element: <Hosts />,
      },
      {
        path: 'sftp',
        element: <Sftp />,
      },
      {
        path: 'command',
        element: <Command />,
      },
      {
        path: 'discovery',
        element: <Discovery />,
      },
      {
        path: 'minio',
        element: <Minio />,
      },
    ],
  },
];

export default routes;