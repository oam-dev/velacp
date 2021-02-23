// prevent Ant design style from being overridden
import 'antd/dist/antd.css';

import React, { useEffect, useState } from 'react';

import { Alert, Button, Card, Form, message, Select, Space } from 'antd';
import { history, useModel } from 'umi';

import { CloseOutlined, SaveOutlined } from '@ant-design/icons';
import { FooterToolbar, PageContainer } from '@ant-design/pro-layout';

import ServiceForm from './components/ServiceForm';
import { ProFormText } from '@ant-design/pro-form';

export default (): React.ReactNode => {
  const [services, setServices] = useState<{ [key: string]: any }>({});

  const [errorFields, setErrorFields] = useState<{ [key: string]: any }>({});
  const [showError, setShowError] = useState<boolean>(false);
  const [environments, setEnvironments] = useState<API.EnvironmentType[]>([]);
  const [caps, setCaps] = useState<API.CapabilityType[]>([]);

  const { addApplication } = useModel('useApplications');
  const { listEnvironments, listCapabilities } = useModel('useEnvironments');

  useEffect(() => {
    listEnvironments().then((r) => setEnvironments(r));
  }, [environments]);

  const saveApp = async (app: API.ApplicationType) => {
    const hide = message.loading('正在添加');
    try {
      await addApplication(app);
      hide();
      message.success('添加成功，即将刷新');
      history.push('/applications');
    } catch (error) {
      hide();
      message.error('添加失败，请重试');
    }
  };

  return (
    <PageContainer>
      {showError ? (
        <Alert
          showIcon
          type="error"
          message="The following fields failed validation:"
          description={
            <ul>
              {Object.keys(errorFields).map((f) =>
                errorFields[f] == null
                  ? null
                  : Object.entries(errorFields[f]).map((ff) => {
                      return (
                        <li key={ff[0]}>
                          - {f}.{ff[0]}
                        </li>
                      );
                    }),
              )}
            </ul>
          }
        />
      ) : null}
      <Form
        labelCol={{ span: 4 }}
        onFinishFailed={({ errorFields: ef }) => {
          setShowError(true);
          const appErrorFields = {};
          ef.forEach((f) => {
            appErrorFields[f.name.join('.')] = f.errors.join(',');
          });
          setErrorFields({ ...errorFields, app: appErrorFields });
        }}
        onFinish={(values) => {
          delete errorFields['app.name'];
          Object.keys(errorFields).forEach((f) => {
            if (errorFields[f] == null) {
              delete errorFields[f];
            }
          });
          if (Object.keys(errorFields).length > 0) {
            setShowError(true);
            return;
          }
          saveApp({ ...values, services });
        }}
      >
        <Space direction="vertical" style={{ width: '100%' }}>
          <Card title="Basic">
            <ProFormText
              width="md"
              name="name"
              label="Application Name"
              rules={[{ required: true, max: 200 }]}
            />
            <ProFormText width="md" name="desc" label="Description" placeholder="请输入名称" />

            <Form.Item name="env" label="Choose environment" required>
              <Select
                onSelect={(e) => {
                  const env = e.toString();
                  listCapabilities(env).then((r) => setCaps(r));
                }}
                placeholder={'Select an environment'}
                optionLabelProp="value"
                style={{ width: '' }}
              >
                {environments.map((item) => (
                  <Select.Option key={item.name} value={item.name}>
                    {item.name}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Card>

          <Card title="Services">
            <ServiceForm
              onChange={(value) => {
                const servicesObj = {};
                value.forEach((service) => {
                  const { name, type, data, traits } = service;
                  if (name == null || type == null) {
                    return;
                  }
                  const serviceObj: any = { type };
                  servicesObj[name] = serviceObj;

                  if (data != null) {
                    Object.keys(data).forEach((k) => {
                      serviceObj[k] = data[k];
                    });
                  }
                  if (traits != null) {
                    Object.keys(traits).forEach((k) => {
                      serviceObj[k] = traits[k];
                    });
                  }
                });
                // if (Object.keys(servicesObj).length < value.length) {
                //   setErrorFields({ ...errorFields, services: { name: '' } });
                //   setShowError(true);
                // } else {
                //   if (errorFields.services?.name != null) {
                //     delete errorFields.services.name;
                //     setErrorFields({ ...errorFields });
                //   }
                //   setShowError(false);
                // }
                setServices(servicesObj);
              }}
              onValidate={(errors) => {
                setErrorFields(errors);
              }}
              caps={caps}
            />
          </Card>
        </Space>
        <FooterToolbar>
          <Button icon={<CloseOutlined />} onClick={() => history.push('/applications')}>
            Cancel
          </Button>
          <Button type="primary" icon={<SaveOutlined />} htmlType="submit">
            Save
          </Button>
        </FooterToolbar>
      </Form>
    </PageContainer>
  );
};