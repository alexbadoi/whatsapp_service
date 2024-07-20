import React, { useState } from 'react';
import { AutoComplete, DatePicker, Button, InputNumber, Select, Input, message, Card, Flex, Menu, Checkbox, Dropdown } from 'antd';
import { SearchOutlined, DownOutlined } from '@ant-design/icons';
import useDebounce from '../../debounce';
import { fetchLocations, parseLocations } from '../../api/location';
import fetchHotelData from '../../api/hotels';

const { RangePicker } = DatePicker;

const HotelSearch = ({ setHotels }) => {
  const [guests, setGuests] = useState({ rooms: 1, adults: 1, children: 0 });
  const [childrenAges, setChildrenAges] = useState([]);
  const [showGuestSelection, setShowGuestSelection] = useState(false);
  const [location, setLocation] = useState('');
  const [dates, setDates] = useState([]);
  const [errors, setErrors] = useState({});
  const [options, setOptions] = useState([]);
  const [selectedLocation, setSelectedLocation] = useState(null);
  const [filters, setFilters] = useState({
    starRating: [],
    guestRating: [],
    propertyFacilities: [],
    bedrooms: [],
    paymentOptions: [],
    popularWithFamilies: false
  });

  const handleGuestChange = (type, value) => {
    setGuests(prev => {
      const newGuests = { ...prev, [type]: value };
      if (type === 'children') {
        adjustChildrenAges(value);
      }
      return newGuests;
    });
  };

  const handleFilterChange = (filterType, value) => {
    setFilters(prev => ({ ...prev, [filterType]: value }));
  };

  const adjustChildrenAges = (newChildrenCount) => {
    setChildrenAges(prevAges => {
      if (newChildrenCount > prevAges.length) {
        return [...prevAges, ...Array(newChildrenCount - prevAges.length).fill(0)];
      } else {
        return prevAges.slice(0, newChildrenCount);
      }
    });
  };

  const handleChildAgeChange = (index, value) => {
    setChildrenAges(prevAges => {
      const newAges = [...prevAges];
      newAges[index] = value;
      return newAges;
    });
  };

  const fetchLocationsDebounced = useDebounce(async (value) => {
    if (value) {
      const data = await fetchLocations(value);
      const parsedOptions = parseLocations(data);
      setOptions(parsedOptions);
    } else {
      setOptions([]);
    }
  }, 1000);

  const handleLocationChange = (value) => {
    setLocation(value);
    setSelectedLocation(null);
    fetchLocationsDebounced(value);
  };
  const handleLocationSelect = (value) => {
    const selectedOption = options.find(option => option.value == value);
    setLocation(selectedOption.label);
    setSelectedLocation(value);
  };

  const validateInputs = () => {
    const newErrors = {};
    if (!selectedLocation) {
      newErrors.location = 'Please select a location from the list';
    }
    if (!dates || dates.length !== 2) {
      newErrors.dates = 'Please select check-in and check-out dates';
    }
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSearch = async () => {
    if (validateInputs()) {
      var from = new Date(dates[0]).toISOString().split('T')[0];
      var to = new Date(dates[1]).toISOString().split('T')[0];
      const data = await fetchHotelData(selectedLocation, from, to, guests.adults, childrenAges, guests.rooms, filters);

      if (data != null) {
        setHotels(data, guests, [from, to], childrenAges, {id: selectedLocation, name: location});
        return;
      }
    }
    message.error('Please correct the errors before searching');
  };

  const filterMenu = (
    <Menu>
      <Menu.ItemGroup title="Star rating">
        <Menu.Item key="5star">
          <Checkbox onChange={(e) => handleFilterChange('starRating', e.target.checked ? [...filters.starRating, 5] : filters.starRating.filter(r => r !== 5))}>
            5 stars
          </Checkbox>
        </Menu.Item>
        <Menu.Item key="4star">
          <Checkbox onChange={(e) => handleFilterChange('starRating', e.target.checked ? [...filters.starRating, 4] : filters.starRating.filter(r => r !== 4))}>
            4 stars
          </Checkbox>
        </Menu.Item>
        <Menu.Item key="norating">
          <Checkbox onChange={(e) => handleFilterChange('starRating', e.target.checked ? [...filters.starRating, 1] : filters.starRating.filter(r => r !== 1))}>
            No rating
          </Checkbox>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.Divider />
      <Menu.ItemGroup title="Guest rating">
        <Menu.Item key="exceptional">
          <Checkbox onChange={(e) => handleFilterChange('guestRating', e.target.checked ? [...filters.guestRating, '9'] : filters.guestRating.filter(r => r !== '9'))}>
            9+ Exceptional
          </Checkbox>
        </Menu.Item>
        <Menu.Item key="excellent">
          <Checkbox onChange={(e) => handleFilterChange('guestRating', e.target.checked ? [...filters.guestRating, '8'] : filters.guestRating.filter(r => r !== '8'))}>
            8+ Excellent
          </Checkbox>
        </Menu.Item>
        <Menu.Item key="verygood">
          <Checkbox onChange={(e) => handleFilterChange('guestRating', e.target.checked ? [...filters.guestRating, '7'] : filters.guestRating.filter(r => r !== '7'))}>
            7+ Very good
          </Checkbox>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.Divider />
      <Menu.ItemGroup title="Property facilities">
        <Menu.Item key="swimmingpool">
          <Checkbox onChange={(e) => handleFilterChange('propertyFacilities', e.target.checked ? [...filters.propertyFacilities, '93'] : filters.propertyFacilities.filter(f => f !== '93'))}>
            Swimming pool
          </Checkbox>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.Divider />
      <Menu.ItemGroup title="Number of bedrooms">
        <Menu.Item key="1bedroom">
          <Checkbox onChange={(e) => handleFilterChange('bedrooms', e.target.checked ? [...filters.bedrooms, 1] : filters.bedrooms.filter(b => b !== 1))}>
            1 bedroom
          </Checkbox>
        </Menu.Item>
        <Menu.Item key="2bedrooms">
          <Checkbox onChange={(e) => handleFilterChange('bedrooms', e.target.checked ? [...filters.bedrooms, 2] : filters.bedrooms.filter(b => b !== 2))}>
            2 bedrooms
          </Checkbox>
        </Menu.Item>
        <Menu.Item key="3bedrooms">
          <Checkbox onChange={(e) => handleFilterChange('bedrooms', e.target.checked ? [...filters.bedrooms, 3] : filters.bedrooms.filter(b => b !== 3))}>
            3 bedrooms
          </Checkbox>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.Divider />
      <Menu.ItemGroup title="Payment options">
        <Menu.Item key="freecancellation">
          <Checkbox onChange={(e) => handleFilterChange('paymentOptions', e.target.checked ? [...filters.paymentOptions, '49499'] : filters.paymentOptions.filter(p => p !== '49499'))}>
            Free cancellation
          </Checkbox>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.Divider />
      <Menu.ItemGroup title="Popular with families">
        <Menu.Item key="kidsstayforFree">
          <Checkbox onChange={(e) => handleFilterChange('popularWithFamilies', e.target.checked)}>
            Kids stay for free
          </Checkbox>
        </Menu.Item>
      </Menu.ItemGroup>
    </Menu>
  );

  return (
    <Card style={{ width: '100%', margin: '0 auto' }}>
      <Flex justify='space-between'>
        <AutoComplete
          options={options}
          style={{ width: '25%', margin: '0 1rem' }}
          value={location}
          onChange={handleLocationChange}
          onSelect={handleLocationSelect}
          status={errors.location ? 'error' : ''}
        >
          <Input.Search
            placeholder="Search for a location"
            enterButton={<SearchOutlined />}
          />
        </AutoComplete>
        <RangePicker
          value={dates}
          onChange={setDates}
          status={errors.dates ? 'error' : ''}
          style={{ margin: '0 1rem' }}
        />
        <div style={{ position: 'relative' }}>
          <Button
            style={{ minWidth: '200px' }}
            onClick={() => setShowGuestSelection(!showGuestSelection)}
          >
            {`${guests.rooms} room, ${guests.adults} adults, ${guests.children} child${guests.children !== 1 ? 'ren' : ''}`}
          </Button>
          {showGuestSelection && (
            <div style={{
              position: 'absolute',
              top: '100%',
              right: 0,
              marginTop: '5px',
              border: '1px solid #d9d9d9',
              padding: '10px',
              backgroundColor: 'white',
              zIndex: 1000,
              width: '300px'
            }}>
              <div style={{ marginBottom: '10px' }}>
                <span style={{ marginRight: '10px', display: 'inline-block', width: '120px' }}>Room</span>
                <InputNumber min={1} value={guests.rooms} onChange={(value) => handleGuestChange('rooms', value)} />
              </div>
              <div style={{ marginBottom: '10px' }}>
                <span style={{ marginRight: '10px', display: 'inline-block', width: '120px' }}>Adults</span>
                <InputNumber min={1} value={guests.adults} onChange={(value) => handleGuestChange('adults', value)} />
              </div>
              <div style={{ marginBottom: '10px' }}>
                <span style={{ marginRight: '10px', display: 'inline-block', width: '120px' }}>Children</span>
                <InputNumber min={0} value={guests.children} onChange={(value) => handleGuestChange('children', value)} />
              </div>
              {childrenAges.map((age, index) => (
                <div key={index} style={{ marginBottom: '10px' }}>
                  <span style={{ marginRight: '10px', display: 'inline-block', width: '120px' }}>Age of Child {index + 1}</span>
                  <Select
                    style={{ width: 120 }}
                    value={age}
                    onChange={(value) => handleChildAgeChange(index, value)}
                  >
                    {[...Array(18)].map((_, i) => (
                      <Select.Option key={i} value={i}>{i}</Select.Option>
                    ))}
                  </Select>
                </div>
              ))}
            </div>
          )}
        </div>
        <Dropdown overlay={filterMenu} trigger={['click']}>
          <Button style={{ width: 100, margin: '0 1rem' }}>
            Filters <DownOutlined />
          </Button>
        </Dropdown>
        <Button type="primary" icon={<SearchOutlined />} onClick={handleSearch}>
          SEARCH
        </Button>
      </Flex>
    </Card>
  );
};

export default HotelSearch;