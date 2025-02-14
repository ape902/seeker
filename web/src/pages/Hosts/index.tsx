import React, { useState, useEffect } from 'react';
import { Table, Card, Button, Input, Space, Modal, Form, message, Popconfirm, Tooltip, Tag, Select, Drawer, InputNumber } from 'antd';
import { PlusOutlined, SearchOutlined } from '@ant-design/icons';
import axios from 'axios';
import '../../styles/common.css';

interface Host {
  id: number;
  ip: string;
  port: number;
  os: string;
  label?: { [key: string]: string };
}

interface HostForm {
  ip: string;
  port: number;
  os: string;
  label?: { [key: string]: string };
}

const Hosts: React.FC = () => {
  const [hosts, setHosts] = useState<Host[]>([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [current, setCurrent] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [searchText, setSearchText] = useState('');
  const [modalVisible, setModalVisible] = useState(false);
  const [modalTitle, setModalTitle] = useState('新增主机');
  const [form] = Form.useForm();
  const [editingId, setEditingId] = useState<number | null>(null);
  const [inputValue, setInputValue] = useState('');
  const [labels, setLabels] = useState<{ [key: string]: string }>({}); 
  const [authMode, setAuthMode] = useState<number>(1);

  const fetchHosts = async () => {
    try {
      setLoading(true);
      const response = await axios.get(`/api/v1/cmdb/list?index=${current}&size=${pageSize}`);
      setHosts(response.data.data.records);
      setTotal(response.data.data.total);
    } catch (error) {
      message.error('获取主机列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchHosts();
  }, [current, pageSize, searchText]);

  const handleAdd = () => {
    form.resetFields();
    setModalTitle('新增主机');
    setEditingId(null);
    setLabels({});
    setAuthMode(2);
    form.setFieldValue('auth_mode', 2);
    setModalVisible(true);
  };

  const handleEdit = async (record: Host) => {
    try {
      const response = await axios.get(`/api/v1/cmdb/info?id=${record.id}`);
      const hostData = response.data.data;
      form.setFieldsValue({
        ip: hostData.ip,
        port: hostData.port,
        os: hostData.os
      });
      setLabels(hostData.label || {});
      setModalTitle('编辑主机');
      setEditingId(record.id);
      setModalVisible(true);
    } catch (error) {
      message.error('获取主机数据失败');
    }
  };

  const validateLabel = (input: string) => {
    const parts = input.split('=');
    return parts.length === 2 && parts[0].trim() && parts[1].trim();
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleInputConfirm = () => {
    if (!inputValue) return;

    if (validateLabel(inputValue)) {
      const [key, value] = inputValue.split('=').map(part => part.trim());
      const newLabels = { ...labels, [key]: value };
      setLabels(newLabels);
      form.setFieldValue('label', newLabels);
      setInputValue('');
    } else {
      message.error('标签格式不正确，请使用 key=value 格式');
    }
  };

  const handleRemoveTag = (key: string) => {
    const newLabels = { ...labels };
    delete newLabels[key];
    setLabels(newLabels);
    form.setFieldValue('label', newLabels);
  };

  const handleDelete = async (id: number) => {
    try {
      await axios.post('/api/v1/cmdb/delete', { ids: [id] });
      message.success('删除成功');
      fetchHosts();
    } catch (error) {
      message.error('删除失败');
    }
  };

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      const submitValues = { ...values, label: labels };

      if (editingId) {
        // 编辑时只提交必要字段
        const updateData = {
          id: editingId,
          ip: submitValues.ip,
          port: submitValues.port,
          os: submitValues.os,
          label: submitValues.label
        };
        await axios.post('/api/v1/cmdb/update', [updateData]);
        message.success('更新成功');
      } else {
        await axios.post('/api/v1/cmdb/create', [submitValues]);
        message.success('创建成功');
      }
      setModalVisible(false);
      fetchHosts();
    } catch (error) {
      message.error('操作失败');
    }
  };

  const columns = [
    {
      title: 'IP地址',
      dataIndex: 'ip',
      key: 'ip',
      render: (_: string, record: Host) => {
        const labels = record.label || {};
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
              <span style={{ fontWeight: 500, color: '#262626' }}>{record.ip}</span>
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
      title: '端口',
      dataIndex: 'port',
      key: 'port',
      render: (port: number) => (
        <div className="status-tag" style={{ backgroundColor: 'rgba(82,196,26,0.1)', color: '#52c41a' }}>
          {port}
        </div>
      ),
    },
    {
      title: '操作系统',
      dataIndex: 'os',
      key: 'os',
      render: (os: string) => (
        <div className="status-tag" style={{ backgroundColor: 'rgba(250,173,20,0.1)', color: '#faad14' }}>
          {os}
        </div>
      ),
    },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: Host) => (
        <Space size="middle">
          <Button
            type="link"
            className="action-button"
            onClick={() => handleEdit(record)}
          >
            编辑
          </Button>
          <Popconfirm
            title="确定要删除该主机吗？"
            onConfirm={() => handleDelete(record.id)}
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
      title="主机管理"
      extra={
        <Space>
          <Input
            placeholder="搜索主机"
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
            新增主机
          </Button>
        </Space>
      }
      className="card-container"
    >
      <Table
        columns={columns}
        dataSource={hosts}
        rowKey="id"
        loading={loading}
        pagination={{
          current,
          pageSize,
          total,
          onChange: (page, size) => {
            setCurrent(page);
            setPageSize(size);
          },
          style: { marginTop: 16 },
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
        <Form
          form={form}
          layout="vertical"
          style={{ marginTop: 16 }}
        >
          <Form.Item
            name="ip"
            label="IP地址"
            rules={[{ required: true, message: '请输入IP地址' }]}
          >
            <Input placeholder="请输入IP地址" />
          </Form.Item>

          <Form.Item
            name="port"
            label="端口"
            rules={[{ required: true, message: '请输入端口号' }, { type: 'number', min: 1, max: 65535, message: '端口号必须在1-65535之间' }]}
          >
            <InputNumber
              style={{ width: '100%' }}
              placeholder="请输入端口号"
              min={1}
              max={65535}
            />
          </Form.Item>

          <Form.Item
            name="os"
            label="操作系统"
            rules={[{ required: true, message: '请输入操作系统' }]}
          >
            <Input placeholder="请输入操作系统" />
          </Form.Item>

          <Form.Item name="label" label="标签">
            <div>
              {Object.entries(labels).map(([key, value]) => (
                <Tag
                  key={key}
                  closable
                  onClose={() => handleRemoveTag(key)}
                  style={{ marginBottom: 8 }}
                >
                  {`${key}=${value}`}
                </Tag>
              ))}
              <Input
                type="text"
                value={inputValue}
                onChange={handleInputChange}
                onPressEnter={handleInputConfirm}
                onBlur={handleInputConfirm}
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

export default Hosts;