import React from 'react';
import { Navigate, RouteObject } from 'react-router-dom';
import BasicLayout from '@/layouts/BasicLayout';

// 错误边界组件
class ErrorBoundary extends React.Component<{ children: React.ReactNode }, { hasError: boolean }> {
  constructor(props: { children: React.ReactNode }) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError() {
    return { hasError: true };
  }

  render() {
    if (this.state.hasError) {
      return <div>页面加载失败，请刷新重试</div>;
    }
    return this.props.children;
  }
}

// 默认重定向组件
const DefaultRedirect = () => {
  const savedSelectedKeys = localStorage.getItem('selectedKeys');
  const defaultPath = savedSelectedKeys ? JSON.parse(savedSelectedKeys)[0] : '/users';
  return <Navigate to={defaultPath} replace />;
};

// 懒加载页面组件
const Login = React.lazy(() => import('@/pages/Login'));
const Users = React.lazy(() => import('@/pages/Users'));
const Hosts = React.lazy(() => import('@/pages/Hosts'));
const Sftp = React.lazy(() => import('@/pages/Sftp'));
const Command = React.lazy(() => import('@/pages/Command'));
const Discovery = React.lazy(() => import('@/pages/Discovery'));
const Minio = React.lazy(() => import('@/pages/Minio'));

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
    element: (
      <ErrorBoundary>
        <Login />
      </ErrorBoundary>
    ),
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
        element: <DefaultRedirect />,
      },
      {
        path: 'users',
        element: (
          <ErrorBoundary>
            <Users />
          </ErrorBoundary>
        ),
      },
      {
        path: 'hosts',
        element: (
          <ErrorBoundary>
            <Hosts />
          </ErrorBoundary>
        ),
      },
      {
        path: 'sftp',
        element: (
          <ErrorBoundary>
            <Sftp />
          </ErrorBoundary>
        ),
      },
      {
        path: 'command',
        element: (
          <ErrorBoundary>
            <Command />
          </ErrorBoundary>
        ),
      },
      {
        path: 'discovery',
        element: (
          <ErrorBoundary>
            <Discovery />
          </ErrorBoundary>
        ),
      },
      {
        path: 'minio',
        element: (
          <ErrorBoundary>
            <Minio />
          </ErrorBoundary>
        ),
      },
    ],
  },
];

export default routes;
