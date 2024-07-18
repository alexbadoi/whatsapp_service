import { Flex, Typography, Steps, Button, message } from 'antd';
import React from 'react';
import HotelSearch from './components/inputs/main';
import HotelListing from './components/listing';
import RoomsCard from './components/rooms';
import { SearchOutlined } from '@ant-design/icons';
import fetchRoomData from './api/rooms';
import SaveProposal from './api/save';

const { Text } = Typography;
const { Step } = Steps;

class AgodaContainer extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      hotels: [],
      hotelSelected: null,
      guests: {},
      dates: [],
      rooms: [],
      currentStep: 0,
      selectedRooms: [],
      childrenAge: [],
      selectedRoomsIds: [],
    };
  }

  getRooms = async (p) => {
    this.setState({ hotelSelected: p});
    const dates = this.state.dates;
    var data = await fetchRoomData(p.PropertyID, dates[0], dates[1], this.state.guests, p.Currency);
    if (data) {
      this.setState({ rooms: data.roomGroup });
    }
  };

  handleNext = () => {
    if (this.state.currentStep === 4) {
      SaveProposal(this.state.selectedRooms)
        .then(res => {
          if (res !== null) {
            this.setState({ currentStep: 0, hotels: [], hotelSelected: null, guests: {}, dates: [], rooms: [], selectedRooms: [], childrenAge: [], selectedRoomsIds: [] });
            message.success('Proposal saved successfully');
          }else{
            message.error('Something went wrong, please try again later');
          }
        });
    }
    this.setState(prevState => ({
      currentStep: Math.min(prevState.currentStep + 1, 4)
    }));
  };

  handlePrevious = () => {
    this.setState(prevState => ({
      currentStep: Math.max(prevState.currentStep - 1, 0)
    }));
  };

  canProceedToNextStep = () => {
    const { currentStep, selectedRooms } = this.state;
    return selectedRooms[currentStep] !== undefined || currentStep === 4;
  };

  selectRoom = (room, group) => {
    var tmp = this.state.selectedRooms;
    var ids = this.state.selectedRoomsIds;
    var hotel = this.state.hotelSelected;
    var proposal = {
      stepCount: this.state.currentStep,
      bookinDt: new Date().toISOString().split('T')[0],
      hotelName: hotel.Name,
      agodaPropertyID: hotel.PropertyID,
      checkIn: this.state.dates[0],
      checkOut: this.state.dates[1],
      numberAdults: this.state.guests.adults,
      numberChildren: this.state.guests.children,
      childrenAge: this.state.childrenAge.join(','), 
      roomName: group.name,
      roomNumber: this.state.guests.rooms,
      commonSpace: group.bedType,
      roomSize: group.sizeInfo,
      cancelationPolicyTitle: room.cancelationPolicy,
      cancelationDeadlineDT: room.cancelationPolicyMaxDate,
      totalPricePerNight: room.price,
      currency: hotel.Currency,
    };
    console.log(proposal);
    tmp[this.state.currentStep] = proposal;
    ids[this.state.currentStep] = room.id;
    this.setState({ selectedRooms: tmp, selectedRoomsIds: ids });
  };

  render() {
    return (
      <Flex align='center' vertical>
        <Steps current={this.state.currentStep} style={styles.main}>
            <Step title="First day" />
            <Step title="Second day" />
            <Step title="Third day" />
            <Step title="Forth day" />
            <Step title="Fifth day" />
          </Steps>
        <Flex align="center" style={styles.main} vertical>
          <HotelSearch setHotels={(h, g, d, c) => this.setState({ hotels: h, guests: g, dates: d, childrenAge: c  })} />
          <Flex justify="space-between" style={{ width: '100%', margin: '1rem 0' }}>
            <Button onClick={this.handlePrevious} disabled={this.state.currentStep === 0}>
              Previous
            </Button>
            <Button onClick={this.handleNext} disabled={!this.canProceedToNextStep()}>
              {this.state.currentStep === 4 ? 'Finish' : 'Next'}
            </Button>
          </Flex>
          {this.state.hotels.length > 0 ? (
            <Flex justify='space-between' style={{ width: '100%', marginTop: '1.5rem' }} align='top'>
              <div style={{ width: '29%' }}>
                {this.state.hotels.map((hotel, i) => (
                  <HotelListing key={i} hotel={hotel} 
                    selectProperty={(p) => this.getRooms(p)}
                  />
                ))}
              </div>
              <div style={{ width: '69%' }}>
                {this.state.hotelSelected && this.state.rooms.length > 0 ? (
                  this.state.rooms.map((room, i) => (
                    <RoomsCard key={i} room={room} 
                      inputRoomNumber={this.state.guests.rooms} setSelectedRoom={(r) => this.selectRoom(r, room)}
                      id={this.state.selectedRoomsIds[this.state.currentStep]}
                    />
                  ))
                ) : (
                  <Flex align="center" justify="center" style={styles.noHotels}>
                    <Text>Please select hotel</Text>
                  </Flex>
                )}
              </div>
            </Flex>
          ) : (
            <Flex align="center" justify="center" style={styles.noHotels}>
              <SearchOutlined />
              <Text>Please search for hotels</Text>
            </Flex>
          )}
        </Flex>
      </Flex>
    );
  }
}

export default AgodaContainer;

const styles = {
  main: {
    marginTop: '2rem',
    maxWidth: '70%',
  },
  noHotels: {
    marginTop: '2rem',
  }
};