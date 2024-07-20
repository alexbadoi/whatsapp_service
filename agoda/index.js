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
      locationId: 0,
    };
  }

  getRooms = async (p) => {
    this.setState({ hotelSelected: p });
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
          } else {
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

  calculateTotal = () => {
    var total = 0;
    for (var i = 0; i < this.state.selectedRooms.length; i++) {
      total += this.state.selectedRooms[i].roomPricePerNight * this.state.selectedRooms[i].numberOfRooms;
    }
    return total.toFixed(2);
  }

  selectRoom = (room, group, roomNumber) => {
    var tmp = this.state.selectedRooms;
    var hotel = this.state.hotelSelected;
    var proposal = {
      dayCount: this.state.currentStep,
      agodaHotelLocationID: this.state.locationId.id,
      agodaLocationName: this.state.locationId.name,
      StartDt: new Date(this.state.dates[0]).toISOString(),
      EndDt: new Date(this.state.dates[1]).toISOString(),
      numberOfRooms: roomNumber,
      numberOfAdults: this.state.guests.adults,
      numberOfChildren: this.state.guests.children,
      childAgesPipeDelimited: this.state.childrenAge.join('|'),
      agodaHotelID: hotel.PropertyID,
      agodaHotelName: hotel.Name,
      agodaHotelStars: hotel.Rating,
      agodaHotelRating: hotel.ReviewScore,
      agodaHotelPicsPipeDelimited: hotel.ImageURL.join('|'),
      agodaRoomPicsPipeDelimited: group.images.join('|'),
      sleepsCount: room.maxOccupancy,
      roomPricePerNight: room.price,
      RiskFreeBookingBinary: room.cancelationPolicy.includes('Risk-Free') ? 1 : 0,
      CancelationDeadlineDt: room.cancelationPolicyMaxDate.includes("Invalid") ? null : new Date(room.cancelationPolicyMaxDate).toISOString(),
      includesBreakfastBinary: room.benefits.includes('Breakfast') ? 1 : 0,
      includesLunchBinary: room.benefits.includes('Lunch') ? 1 : 0,
      includesDinnerBinary: room.benefits.includes('Dinner') ? 1 : 0,
      bedType: group.bedType,
      sizeSqM: group.sizeInfo,
    };
    console.log(proposal);
    tmp.push(proposal);
    this.setState({ selectedRooms: tmp });
  };

  render() {
    return (
      <Flex align='center' vertical>
        <Flex style={styles.main} align='center' justify="space-between">
          <Steps current={this.state.currentStep} style={{ width: '85%' }}>
            <Step title="First day" />
            <Step title="Second day" />
            <Step title="Third day" />
            <Step title="Forth day" />
            <Step title="Fifth day" />
          </Steps>
          <Text style={{ fontSize: '1rem', fontWeight: 'bold' }} >Total: {this.calculateTotal()} $</Text>
        </Flex>
        <Flex align="center" style={styles.main} vertical>
          <HotelSearch setHotels={(h, g, d, c, l) => this.setState({ hotels: h, guests: g, dates: d, childrenAge: c, locationId: l })} />
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
              <div style={{ width: '39%' }}>
                {this.state.hotels.map((hotel, i) => (
                  <HotelListing key={i} hotel={hotel}
                    selectProperty={(p) => this.getRooms(p)}
                  />
                ))}
              </div>
              <div style={{ width: '59%' }}>
                {this.state.hotelSelected && this.state.rooms.length > 0 ? (
                  this.state.rooms.map((room, i) => (
                    <RoomsCard key={i} room={room}
                      inputRoomNumber={this.state.guests.rooms} setSelectedRoom={(r, roomNumber) => this.selectRoom(r, room, roomNumber)}
                      shouldDisableButton={(offer) => this.state.selectedRooms.some(r => r.roomId === offer.id && r.stepCount === this.state.currentStep)}
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
    width: '98%',
  },
  noHotels: {
    marginTop: '2rem',
  }
};