import React, { useState, useEffect } from 'react';
import { Table, Card, Button, Input, Space, message, Popconfirm } from 'antd';
import { PlusOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons';
import axios from 'axios';
import '../../styles/common.css';

interface Bucket {
  name: string;
  created_at: string;
}

const Minio: React.FC = () => {
  const [buckets, setBuckets] = useState<Bucket[]>([]);
  const [loading, setLoading] = useState(false);
  const [searchText, setSearchText] = useState('');

  const fetchBuckets = async () => {
    try {
      setLoading(true);
      const response = await axios.get('/api/minio/buckets');
      const filteredBuckets = response.data.items.filter(bucket =>
        bucket.name.toLowerCase().includes(searchText.toLowerCase())
      );
      setBuckets(filteredBuckets);
    } catch (error) {
      message.error('获取存储桶列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBuckets();
  }, [searchText]);

  const handleCreate = async () => {
    const bucketName = prompt('请输入存储桶名称：');
    if (!bucketName) return;

    try {
      await axios.post('/api/minio/buckets', { name: bucketName });
      message.success('创建成功');
      fetchBuckets();
    } catch (error) {
      message.error('创建失败');
    }
  };

  const handleDelete = async (name: string) => {
    try {
      await axios.delete(`/api/minio/buckets/${name}`);
      message.success('删除成功');
      fetchBuckets();
    } catch (error) {
      message.error('删除失败');
    }
  };

  const columns = [
    {
      title: '存储桶名称',
      dataIndex: 'name',
      key: 'name',
      render: (text: string) => (
        <span style={{ fontWeight: 500, color: '#262626' }}>{text}</span>
      ),
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (text: string) => (
        <span style={{ color: '#595959' }}>{text}</span>
      ),
    },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: Bucket) => (
        <Space>
          <Popconfirm
            title="确定要删除该存储桶吗？"
            onConfirm={() => handleDelete(record.name)}
          >
            <Button 
              type="link" 
              danger 
              icon={<DeleteOutlined />}
              className="action-button"
            >
              删除
            </Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <Card
      title="存储桶管理"
      extra={
        <Space>
          <Input
            placeholder="搜索存储桶"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
            prefix={<SearchOutlined style={{ color: '#bfbfbf' }} />}
            className="search-input"
          />
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={handleCreate}
          >
            新增存储桶
          </Button>
        </Space>
      }
      className="card-container"
    >
      <Table
        columns={columns}
        dataSource={buckets}
        rowKey="name"
        loading={loading}
        pagination={false}
        className="table-container"
      />
    </Card>
  );
};

export default Minio;