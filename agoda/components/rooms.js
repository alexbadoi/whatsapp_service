import React, { useEffect } from 'react';
import { Card, List, Tag, Button, Typography, InputNumber, Carousel, Divider, Flex } from 'antd';
import { UserOutlined } from '@ant-design/icons';

const { Title, Text } = Typography;

const Rooms = ({ room, inputRoomNumber, setSelectedRoom, shouldDisableButton }) => {

  const [rooms, setRooms] = React.useState([]);

  useEffect(() => {
    setRooms(new Array(room.rooms.length).fill(inputRoomNumber));
  }, []);

  const startsWithRiskFree = (str) => {
    return str.includes('Risk-free');
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

        <Flex vertical>
          {room.rooms.map((offer, index) => (
            <React.Fragment key={index}>
              <Card>
                <Flex justify="space-between" align="flex-start" style={{ marginBottom: '1rem', }} vertical>
                  <Flex justify='space-between' style={{ width: '100%' }}>
                    <Flex vertical align='flex-start'>
                      <Text style={{ fontSize: '12px' }}>Your price includes:</Text>
                      <Flex wrap>
                        {offer.benefits.map((benefit, index) => (
                          <Text key={index} style={{ fontSize: '14px',  marginRight: '2px',
                            color: benefit == "FULL BOARD" ? 'green' : 'black',
                            fontWeight: benefit == "FULL BOARD" ? 'bolder' : 'normal'
                           }}>{benefit}{index < offer.benefits.length - 1 ? ", " : ""}</Text>
                        ))}
                      </Flex>
                    </Flex>
                    <div>
                      <Tag icon={<UserOutlined />} style={{ fontSize: '16px', padding: '4px 8px', }}>{offer.maxOccupancy}</Tag>
                    </div>
                  </Flex>

                  <Flex justify='space-between' align="center" style={{ width: '100%' }}>
                    <div style={{ width: '69%' }}>
                      <Title level={5} style={{
                        fontSize: '12px',
                        color: startsWithRiskFree(offer.cancelationPolicy) ? 'green' : 'black',
                        fontWeight: startsWithRiskFree(offer.cancelationPolicy) ? 'bolder' : 'normal'
                      }}>{offer.cancelationPolicy}</Title>
                      {!offer.cancelationPolicyMaxDate.includes("Invalid") &&
                        <Text style={{ fontSize: '10px' }}><b>Free until:</b> {offer.cancelationPolicyMaxDate}</Text>
                      }
                    </div>
                    <Flex style={{ width: '29%' }} vertical>
                      <Flex vertical align='flex-end'>
                        <Title level={3} style={{ margin: 0, fontSize: '18px' }}>{offer.price}$</Title>
                        <Text type="secondary" style={{ fontSize: '10px', textAlign: 'right' }}>Per night after taxes and fees</Text>
                      </Flex>
                      <Flex vertical>
                        <Flex align='flex-end' vertical>
                          <InputNumber min={1} value={rooms[index]} onChange={(value) => {var tmp = rooms;tmp[index] = value; setRooms(tmp);}} style={{ maxWidth: '50px' }} />
                          <Button type="primary" size="small" style={{ flex: 1, fontSize: '14px', marginTop: '.5rem' }} 
                            onClick={() => setSelectedRoom(offer, rooms[index])}
                            disabled={shouldDisableButton(offer)}
                            >Add selection</Button>
                        </Flex>
                      </Flex>
                    </Flex>
                  </Flex>
                </Flex>
              </Card>
              {index < room.rooms.length - 1 && <Divider style={{ margin: '12px 0' }} />}
            </React.Fragment>
          ))}
        </Flex>
      </Flex>
    </Card>
  );
};

export default Rooms;