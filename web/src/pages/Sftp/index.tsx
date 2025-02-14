import React, { useState, useEffect } from 'react';
import { Table, Card, Button, Space, Upload, message } from 'antd';
import { UploadOutlined, DownloadOutlined, DeleteOutlined } from '@ant-design/icons';
import type { UploadProps } from 'antd';
import axios from 'axios';
import '../../styles/common.css';

interface FileItem {
  key: string;
  name: string;
  size: number;
  type: string;
  modified_at: string;
}

const Sftp: React.FC = () => {
  const [files, setFiles] = useState<FileItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [currentPath, setCurrentPath] = useState('/');

  const fetchFiles = async () => {
    try {
      setLoading(true);
      const response = await axios.get(`/api/sftp/files?path=${currentPath}`);
      setFiles(response.data.items);
    } catch (error) {
      message.error('获取文件列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchFiles();
  }, [currentPath]);

  const handleDownload = async (file: FileItem) => {
    try {
      const response = await axios.get(`/api/sftp/download?path=${currentPath}/${file.name}`, {
        responseType: 'blob'
      });
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', file.name);
      document.body.appendChild(link);
      link.click();
      link.remove();
      message.success('下载成功');
    } catch (error) {
      message.error('下载文件失败');
    }
  };

  const handleDelete = async (file: FileItem) => {
    try {
      await axios.delete(`/api/sftp/files?path=${currentPath}/${file.name}`);
      message.success('删除成功');
      fetchFiles();
    } catch (error) {
      message.error('删除失败');
    }
  };

  const uploadProps: UploadProps = {
    name: 'file',
    action: `/api/sftp/upload?path=${currentPath}`,
    headers: {
      Authorization: `Bearer ${localStorage.getItem('token')}`,
    },
    onChange(info) {
      if (info.file.status === 'done') {
        message.success(`${info.file.name} 上传成功`);
        fetchFiles();
      } else if (info.file.status === 'error') {
        message.error(`${info.file.name} 上传失败`);
      }
    },
  };

  const columns = [
    {
      title: '文件名',
      dataIndex: 'name',
      key: 'name',
      render: (text: string) => (
        <span style={{ fontWeight: 500, color: '#262626' }}>{text}</span>
      ),
    },
    {
      title: '类型',
      dataIndex: 'type',
      key: 'type',
      render: (text: string) => (
        <span style={{ color: '#595959' }}>{text}</span>
      ),
    },
    {
      title: '大小',
      dataIndex: 'size',
      key: 'size',
      render: (size: number) => {
        const units = ['B', 'KB', 'MB', 'GB'];
        let value = size;
        let unitIndex = 0;
        while (value >= 1024 && unitIndex < units.length - 1) {
          value /= 1024;
          unitIndex++;
        }
        return (
          <span style={{ color: '#595959' }}>
            {`${value.toFixed(2)} ${units[unitIndex]}`}
          </span>
        );
      },
    },
    {
      title: '修改时间',
      dataIndex: 'modified_at',
      key: 'modified_at',
      render: (text: string) => (
        <span style={{ color: '#8c8c8c' }}>{text}</span>
      ),
    },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: FileItem) => (
        <Space>
          <Button
            type="link"
            icon={<DownloadOutlined />}
            onClick={() => handleDownload(record)}
            className="action-button"
          >
            下载
          </Button>
          <Button
            type="link"
            danger
            icon={<DeleteOutlined />}
            onClick={() => handleDelete(record)}
            className="action-button"
          >
            删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <Card className="card-container">
      <Space style={{ marginBottom: 16 }}>
        <Upload {...uploadProps}>
          <Button icon={<UploadOutlined />}>上传文件</Button>
        </Upload>
      </Space>

      <Table
        columns={columns}
        dataSource={files}
        loading={loading}
        pagination={false}
        className="table-container"
      />
    </Card>
  );
};

export default Sftp;