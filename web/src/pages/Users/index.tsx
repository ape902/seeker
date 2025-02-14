import React, { useState, useEffect } from 'react';
import { Table, Card, Button, Input, Space, Modal, Form, message, Popconfirm, Select, Tooltip, Tag, Drawer } from 'antd';
import { PlusOutlined, SearchOutlined } from '@ant-design/icons';
import axios from 'axios';
import '../../styles/common.css';

interface User {
  id: number;
  nick_name: string;
  mobile: string;
  rule: number;
  labels?: { [key: string]: string };
}

interface UserForm {
  nick_name: string;
  mobile: string;
  password?: string;
  rule: number;
  labels?: { [key: string]: string };
}

const Users: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [current, setCurrent] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [searchText, setSearchText] = useState('');
  const [modalVisible, setModalVisible] = useState(false);
  const [modalTitle, setModalTitle] = useState('新增用户');
  const [form] = Form.useForm();
  const [editingId, setEditingId] = useState<number | null>(null);
  const [inputValue, setInputValue] = useState('');

  const fetchUsers = async () => {
    try {
      setLoading(true);
      const response = await axios.get(`/api/v1/user/list?index=${current}&size=${pageSize}&search=${searchText}`);
      setUsers(response.data.data.records);
      setTotal(response.data.data.total);
    } catch (error) {
      message.error('获取用户列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, [current, pageSize, searchText]);

  const handleAdd = () => {
    form.resetFields();
    setModalTitle('新增用户');
    setEditingId(null);
    setModalVisible(true);
  };

  const handleEdit = async (record: User) => {
    try {
      const response = await axios.get(`/api/v1/user/info?id=${record.id}`);
      const userData = response.data.data;
      form.setFieldsValue({
        nick_name: userData.nick_name,
        mobile: userData.mobile,
        rule: userData.rule,
        labels: userData.labels || {}
      });
      setModalTitle('编辑用户');
      setEditingId(record.id);
      setModalVisible(true);
    } catch (error) {
      message.error('获取用户数据失败');
    }
  };

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      const submitValues = { ...values };

      if (editingId) {
        await axios.post('/api/v1/user/modify', [{ ...submitValues, id: editingId }]);
        message.success('更新成功');
      } else {
        await axios.post('/api/v1/user/add', submitValues);
        message.success('创建成功');
      }
      setModalVisible(false);
      fetchUsers();
    } catch (error) {
      message.error('操作失败');
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await axios.post('/api/v1/user/del', { ids: [id] });
      message.success('删除成功');
      fetchUsers();
    } catch (error) {
      message.error('删除失败');
    }
  };

  const columns = [
    {
      title: '用户名',
      dataIndex: 'nick_name',
      key: 'nick_name',
      render: (text: string) => (
        <span style={{ fontWeight: 500, color: '#262626' }}>{text}</span>
      ),
    },
    {
      title: '手机号',
      dataIndex: 'mobile',
      key: 'mobile',
      render: (_: any, record: User) => {
        const labels = record.labels || {};
        return (
          <Tooltip
            title={
              <div className="tooltip-content">
                {Object.entries(labels).length > 0 ? (
                  <div style={{ display: 'flex', flexWrap: 'wrap', gap: '8px' }}>
                    {Object.entries(labels).map(([key, value]) => (
                      <div key={key} className="label-tag">
                        <span className="label-key">{key}:</span>
                        <span className="label-value">{value}</span>
                      </div>
                    ))}
                  </div>
                ) : (
                  <span style={{ color: '#999' }}>暂无标签</span>
                )}
              </div>
            }
            color="#fff"
            placement="right"
            overlayStyle={{ animationDuration: '0.2s' }}
          >
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
              <span style={{ color: '#595959' }}>{record.mobile}</span>
              {Object.keys(labels).length > 0 && (
                <div className="label-count">
                  {Object.keys(labels).length}
                </div>
              )}
            </div>
          </Tooltip>
        );
      },
    },
    {
      title: '角色',
      dataIndex: 'rule',
      key: 'rule',
      render: (rule: number) => (
        <Tag
          color={rule === 1 ? '#52c41a' : '#faad14'}
          className="status-tag"
        >
          {rule === 1 ? '管理员' : '普通用户'}
        </Tag>
      ),
    },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: User) => (
        <Space size="middle">
          <Button
            type="link"
            onClick={() => handleEdit(record)}
            className="action-button"
          >
            编辑
          </Button>
          <Popconfirm
            title="确定要删除该用户吗？"
            onConfirm={() => handleDelete(record.id)}
            okText="确定"
            cancelText="取消"
          >
            <Button
              type="link"
              danger
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
      title="用户管理"
      extra={
        <Space>
          <Input
            placeholder="搜索用户"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
            prefix={<SearchOutlined style={{ color: '#bfbfbf' }} />}
            className="search-input"
          />
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={handleAdd}
          >
            新增用户
          </Button>
        </Space>
      }
      className="card-container"
    >
      <Table
        columns={columns}
        dataSource={users}
        loading={loading}
        rowKey="id"
        pagination={{
          current,
          pageSize,
          total,
          onChange: (page, size) => {
            setCurrent(page);
            setPageSize(size);
          },
          showSizeChanger: true,
          showQuickJumper: true,
          showTotal: (total) => `共 ${total} 条记录`,
          style: { marginTop: '16px' },
        }}
        className="table-container"
      />

      <Drawer
        title={modalTitle}
        open={modalVisible}
        onClose={() => setModalVisible(false)}
        width={480}
        extra={
          <Space>
            <Button onClick={() => setModalVisible(false)}>取消</Button>
            <Button type="primary" onClick={handleSubmit} loading={loading}>
              确定
            </Button>
          </Space>
        }
      >
        <Form form={form} layout="vertical" style={{ marginTop: 16 }}>
          <Form.Item
            name="nick_name"
            label="用户名"
            rules={[{ required: true, message: '请输入用户名' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="mobile"
            label="手机号"
            rules={[{ required: true, message: '请输入手机号' }]}
          >
            <Input />
          </Form.Item>
          {!editingId && (
            <Form.Item
              name="password"
              label="密码"
              rules={[{ required: true, message: '请输入密码' }]}
            >
              <Input.Password />
            </Form.Item>
          )}
          <Form.Item
            name="rule"
            label="角色"
            rules={[{ required: true, message: '请选择角色' }]}
          >
            <Select
              options={[
                { value: 1, label: '管理员' },
                { value: 2, label: '普通用户' }
              ]}
            />
          </Form.Item>
          <Form.Item name="labels" label="标签">
            <div>
              {Object.entries(form.getFieldValue('labels') || {}).map(([key, value]) => (
                <Tag
                  key={key}
                  closable
                  onClose={() => {
                    const currentLabels = { ...form.getFieldValue('labels') };
                    delete currentLabels[key];
                    form.setFieldValue('labels', currentLabels);
                  }}
                  style={{ marginBottom: 8 }}
                >
                  {`${key}=${value}`}
                </Tag>
              ))}
              <Input
                type="text"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                onPressEnter={() => {
                  if (!inputValue) return;
                  const parts = inputValue.split('=');
                  if (parts.length === 2 && parts[0].trim() && parts[1].trim()) {
                    const [key, value] = parts.map(part => part.trim());
                    const currentLabels = { ...form.getFieldValue('labels') || {} };
                    form.setFieldValue('labels', { ...currentLabels, [key]: value });
                    setInputValue('');
                  } else {
                    message.error('标签格式不正确，请使用 key=value 格式');
                  }
                }}
                onBlur={() => {
                  if (!inputValue) return;
                  const parts = inputValue.split('=');
                  if (parts.length === 2 && parts[0].trim() && parts[1].trim()) {
                    const [key, value] = parts.map(part => part.trim());
                    const currentLabels = { ...form.getFieldValue('labels') || {} };
                    form.setFieldValue('labels', { ...currentLabels, [key]: value });
                    setInputValue('');
                  } else {
                    message.error('标签格式不正确，请使用 key=value 格式');
                  }
                }}
                placeholder="输入标签 (格式: key=value)"
                style={{ width: 200 }}
              />
            </div>
          </Form.Item>
        </Form>
      </Drawer>
    </Card>
  );
};

export default Users;