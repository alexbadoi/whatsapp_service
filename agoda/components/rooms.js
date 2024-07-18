import React from 'react';
import { Card, List, Space, Tag, Button, Typography, Tooltip, Carousel, Divider, Flex } from 'antd';
import { UserOutlined, CheckOutlined } from '@ant-design/icons';

const { Title, Text } = Typography;

const Rooms = ({ room, inputRoomNumber, setSelectedRoom, id }) => {
  const sectionStyle = {
    backgroundColor: '#f0f0f0',
    padding: '8px',
    marginBottom: '8px',
    textAlign: 'center',
  };

  const numberStyle = {
    backgroundColor: '#f0f0f0',
    padding: '4px 8px',
    borderRadius: '4px',
    margin: 0,
    marginRight: '1rem',
  };
  return (
    <Card title={room.name} size="small" style={{ width: '100%' }}>
      <Flex>
        <Flex vertical style={{ marginRight: '16px', width: '180px' }}>
          <Carousel arrows infinite={true}>
            {room.images.map((image, index) => (
              <div key={index}>
                <img src={image} alt={`room ${index + 1}`} style={{ height: '200px', objectFit: 'cover', width: '100%' }} />
              </div>
            ))}
          </Carousel>
          <List
            size="small"
            dataSource={[
              `Bed Type: ${room.bedType}`,
              `Size: ${room.sizeInfo}`,
              ...room.features
            ]}
            renderItem={(item) => <List.Item style={{ fontSize: '10px' }}>{item}</List.Item>}
          />
        </Flex>

        <Flex vertical style={{ flex: 1 }}>
          {room.rooms.map((offer, index) => (
            <React.Fragment key={index}>
              <Flex justify="space-between" align="flex-start" style={{ marginBottom: '1rem' }}>
                <Flex vertical style={{ flex: 1 }}>
                  <Title level={5} style={{ ...sectionStyle, fontSize: '14px' }}>Benefits</Title>
                  <Text style={{ fontSize: '10px' }}>Your price includes:</Text>
                  <List
                    size="small"
                    dataSource={[
                      ...offer.benefits,
                    ]}
                    renderItem={(item) => (
                      <List.Item style={{ padding: '4px 0' }}>
                        <Space>
                          <CheckOutlined style={{ color: 'green', fontSize: '10px' }} />
                          <Text style={{ fontSize: '10px' }}>{item}</Text>
                        </Space>
                      </List.Item>
                    )}
                  />
                </Flex>

                <Flex vertical style={{ width: '80px', alignItems: 'center' }}>
                  <Title level={5} style={{ ...sectionStyle, fontSize: '14px' }}>Sleeps</Title>
                  <Tooltip title={`Sleeps ${offer.maxOccupancy}`}>
                    <Tag icon={<UserOutlined />} style={{ fontSize: '16px', padding: '4px 8px' }}>{offer.maxOccupancy}</Tag>
                  </Tooltip>
                </Flex>

                <Flex vertical style={{ width: '120px', alignItems: 'center' }}>
                  <Title level={5} style={{ ...sectionStyle, fontSize: '14px' }}>Price per night</Title>
                  <Title level={3} style={{ margin: 0, fontSize: '18px' }}>{offer.price}$</Title>
                  <Text type="secondary" style={{ fontSize: '8px' }}>Per night after taxes and fees</Text>
                </Flex>

                <Flex vertical style={{ width: '150px' }}>
                  <Title level={5} style={{ ...sectionStyle, fontSize: '14px' }}>Most booked</Title>
                  <Flex align="center">
                    <Title level={4} style={{ ...numberStyle, fontSize: '16px' }}>{inputRoomNumber}</Title>
                    <Button type="primary" size="small" style={{ flex: 1, fontSize: '12px' }} onClick={()=> setSelectedRoom(offer)} disabled={(offer.id == id)}>Book now</Button>
                  </Flex>
                  <Title level={5} style={{ fontSize: '12px' }}>{offer.cancelationPolicy}</Title>
                  <Text style={{ fontSize: '10px' }}><b>Free until:</b> {offer.cancelationPolicyMaxDate}</Text>
                </Flex>
              </Flex>
              {index < room.rooms.length - 1 && <Divider style={{ margin: '12px 0' }} />}
            </React.Fragment>
          ))}
        </Flex>
      </Flex>
    </Card>
  );
};

export default Rooms;