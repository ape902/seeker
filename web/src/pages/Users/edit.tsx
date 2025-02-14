import React, { useState } from 'react';
import { Form, Input, Button, Tag, message } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

interface UserEditProps {
  initialValues?: {
    id?: number;
    mobile?: string;
    nick_name?: string;
    password?: string;
    labels?: { [key: string]: string };
  };
  onSubmit: (values: any) => void;
}

const UserEdit: React.FC<UserEditProps> = ({ initialValues, onSubmit }) => {
  const [form] = Form.useForm();
  const [inputValue, setInputValue] = useState('');
  const [labels, setLabels] = useState<{ [key: string]: string }>(initialValues?.labels || {});

  // 验证标签格式
  const validateLabel = (input: string) => {
    const parts = input.split('=');
    return parts.length === 2 && parts[0].trim() && parts[1].trim();
  };

  // 处理标签输入
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  // 处理标签输入确认
  const handleInputConfirm = () => {
    if (!inputValue) return;

    if (validateLabel(inputValue)) {
      const [key, value] = inputValue.split('=').map(part => part.trim());
      const newLabels = { ...labels, [key]: value };
      setLabels(newLabels);
      form.setFieldValue('labels', newLabels);
      setInputValue('');
    } else {
      message.error('标签格式不正确，请使用 key=value 格式');
    }
  };

  // 删除标签
  const handleRemoveTag = (key: string) => {
    const newLabels = { ...labels };
    delete newLabels[key];
    setLabels(newLabels);
    form.setFieldValue('labels', newLabels);
  };

  // 处理表单提交
  const handleSubmit = (values: any) => {
    const submitValues = {
      ...values,
      labels: labels
    };
    onSubmit(submitValues);
  };

  return (
    <Form
      form={form}
      initialValues={initialValues}
      onFinish={handleSubmit}
      layout="vertical"
    >
      <Form.Item name="mobile" label="手机号码" rules={[{ required: true }]}>
        <Input />
      </Form.Item>

      <Form.Item name="nick_name" label="昵称" rules={[{ required: true }]}>
        <Input />
      </Form.Item>

      <Form.Item name="password" label="密码">
        <Input.Password placeholder="不修改请留空" />
      </Form.Item>

      <Form.Item name="labels" label="标签">
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

      <Form.Item>
        <Button type="primary" htmlType="submit">
          保存
        </Button>
      </Form.Item>
    </Form>
  );
};

export default UserEdit;