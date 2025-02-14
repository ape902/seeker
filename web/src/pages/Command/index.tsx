import React, { useState } from 'react';
import { Card, Input, Button, Select, Space, message } from 'antd';
import { SendOutlined } from '@ant-design/icons';
import axios from 'axios';
import '../../styles/common.css';

interface CommandResult {
  output: string;
  error: string;
  exit_code: number;
}

const Command: React.FC = () => {
  const [command, setCommand] = useState('');
  const [selectedHosts, setSelectedHosts] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<CommandResult | null>(null);

  const handleExecute = async () => {
    if (!command.trim()) {
      message.warning('请输入要执行的命令');
      return;
    }

    if (selectedHosts.length === 0) {
      message.warning('请选择要执行命令的主机');
      return;
    }

    try {
      setLoading(true);
      const response = await axios.post('/api/command/execute', {
        command,
        hosts: selectedHosts,
      });
      setResult(response.data);
      message.success('命令执行成功');
    } catch (error) {
      message.error('命令执行失败');
      setResult(null);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card className="card-container">
      <Space direction="vertical" style={{ width: '100%' }} size="large">
        <Select
          mode="multiple"
          style={{ width: '100%' }}
          placeholder="选择要执行命令的主机"
          onChange={(values) => setSelectedHosts(values)}
          options={[]}
          loading={loading}
        />

        <Input.TextArea
          placeholder="输入要执行的命令"
          value={command}
          onChange={(e) => setCommand(e.target.value)}
          autoSize={{ minRows: 3 }}
        />

        <Button
          type="primary"
          icon={<SendOutlined />}
          onClick={handleExecute}
          loading={loading}
          className="action-button"
        >
          执行命令
        </Button>

        {result && (
          <Card title="执行结果" size="small" className="table-container">
            <Input.TextArea
              value={result.output || result.error}
              readOnly
              autoSize={{ minRows: 5, maxRows: 15 }}
              style={{
                backgroundColor: '#f5f5f5',
                color: result.exit_code === 0 ? '#000' : '#ff4d4f',
              }}
            />
            <div style={{ marginTop: 8 }}>
              退出码: {result.exit_code}
            </div>
          </Card>
        )}
      </Space>
    </Card>
  );
};

export default Command;