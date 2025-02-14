import React, { useState, useEffect } from 'react';
import { Table, Card, Tag, Space, Button, message } from 'antd';
import { SyncOutlined } from '@ant-design/icons';
import axios from 'axios';
import '../../styles/common.css';

interface Service {
  id: number;
  name: string;
  host: string;
  port: number;
  status: string;
  health: string;
  last_check: string;
}

const Discovery: React.FC = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchServices = async () => {
    try {
      setLoading(true);
      const response = await axios.get('/api/discovery/services');
      setServices(response.data.items);
    } catch (error) {
      message.error('获取服务列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchServices();
    const timer = setInterval(fetchServices, 30000);
    return () => clearInterval(timer);
  }, []);

  const handleRefresh = () => {
    fetchServices();
  };

  const getStatusTag = (status: string) => {
    const colorMap: { [key: string]: string } = {
      running: '#52c41a',
      stopped: '#ff4d4f',
      unknown: '#d9d9d9',
    };
    return (
      <Tag
        color={colorMap[status] || '#d9d9d9'}
        className="status-tag"
      >
        {status.toUpperCase()}
      </Tag>
    );
  };

  const getHealthTag = (health: string) => {
    const colorMap: { [key: string]: string } = {
      healthy: '#52c41a',
      unhealthy: '#ff4d4f',
      warning: '#faad14',
    };
    return (
      <Tag
        color={colorMap[health] || '#d9d9d9'}
        className="status-tag"
      >
        {health.toUpperCase()}
      </Tag>
    );
  };

  const columns = [
    {
      title: '服务名称',
      dataIndex: 'name',
      key: 'name',
      render: (text: string) => (
        <span style={{ fontWeight: 500, color: '#262626' }}>{text}</span>
      ),
    },
    {
      title: '主机地址',
      dataIndex: 'host',
      key: 'host',
      render: (text: string) => (
        <span style={{ color: '#595959' }}>{text}</span>
      ),
    },
    {
      title: '端口',
      dataIndex: 'port',
      key: 'port',
      render: (text: number) => (
        <span style={{ color: '#595959' }}>{text}</span>
      ),
    },
    {
      title: '运行状态',
      dataIndex: 'status',
      key: 'status',
      render: (status: string) => getStatusTag(status),
    },
    {
      title: '健康状态',
      dataIndex: 'health',
      key: 'health',
      render: (health: string) => getHealthTag(health),
    },
    {
      title: '最后检查时间',
      dataIndex: 'last_check',
      key: 'last_check',
      render: (text: string) => (
        <span style={{ color: '#8c8c8c' }}>{text}</span>
      ),
    },
  ];

  return (
    <Card className="card-container">
      <Space style={{ marginBottom: 16 }}>
        <Button
          type="primary"
          icon={<SyncOutlined />}
          onClick={handleRefresh}
          loading={loading}
          className="action-button"
        >
          刷新
        </Button>
      </Space>

      <Table
        columns={columns}
        dataSource={services}
        rowKey="id"
        loading={loading}
        pagination={{
          showSizeChanger: true,
          showQuickJumper: true,
          showTotal: (total) => `共 ${total} 条`,
          style: { marginTop: '16px' },
        }}
        className="table-container"
      />
    </Card>
  );
};

export default Discovery;