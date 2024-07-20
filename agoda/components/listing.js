import React from 'react';
import { Card, Rate, Tag, Typography, Button, Image, Carousel, Flex } from 'antd';

const { Text, Title } = Typography;

const HotelListing = ({ hotel, selectProperty }) => {
  return (
    <Card
      hoverable
      style={{ width: '100%', marginBottom: '1rem' }}
    >
      <Flex align='flex-start'>
        <div style={{ width: '50%' }}>
          <Carousel
            autoplay
            arrows infinite={true}
            style={{ width: '100%' }}
          >
            {hotel.ImageURL.map((image, index) => (
              <Image
                key={index}
                src={image}
                width={260}
                height={180}
              />
            ))}
          </Carousel>
          <Flex wrap>
              {hotel.Amenities && hotel.Amenities.map((amenity, index) => (
                <Tag key={index} color="blue" style={{ fontSize: '10px', marginTop: '.5rem' }}>{amenity}</Tag>
              ))}
            </Flex>
        </div>
        <div style={{ width: '50%', marginLeft: '1rem' }}>
          <Flex vertical justify='start' >
            <Title level={4} style={{ margin: 0, fontSize: '18px' }}>{hotel.Name}</Title>
            <Rate disabled defaultValue={5} value={hotel.Rating} style={{ fontSize: 10 }} />
            <Text type="secondary" style={{ fontSize: '12px' }}>{hotel.Location}</Text>
            <Flex justify='space-between' style={{ marginTop: '1rem' }}>
              <Flex vertical align='start'>
                <div style={{ backgroundColor: '#1890ff', color: 'white', padding: '2px 8px', borderRadius: '12px' }}>
                  <Text strong style={{ color: 'white', fontSize: '12px' }}>{hotel.ReviewScore}</Text>
                </div>
                <Text type="secondary" style={{ fontSize: '10px' }}>{hotel.TotalReviews} reviews</Text>
              </Flex>
              {hotel.PricePerNight != 0 &&
                <Flex vertical align='end'>
                  <Title level={5} style={{ margin: 0, fontSize: '14px' }}>Price per night</Title>
                  <Title level={5} style={{ color: '#ff4d4f', margin: 0, fontSize: '14px' }}>
                    {hotel.PricePerNight}$
                  </Title>
                </Flex>
              }
            </Flex>
            <Button type="primary" style={{ width: '100%', marginTop: '1rem', fontSize: '12px' }} onClick={() => selectProperty(hotel)} >See Rooms</Button>
          </Flex>
        </div>
      </Flex>
    </Card>
  );
};

export default HotelListing;