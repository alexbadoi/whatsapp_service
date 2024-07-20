const buildRoomApiUrl = (id, checkIn, checkOut, adult, children, room, currency) => {
    const baseUrl = '/api/rooms';
    const params = new URLSearchParams({
        id,
        checkIn,
        checkOut,
        adult,
        room,
        currency
    });
    if (children.length > 0) {
        if (children.length === 1) {
          params.append('childrens', children[0]);
        } else {
          params.append('childrens', children.join(','));
        }
      }
    return `${baseUrl}?${params.toString()}`;
};

const fetchRoomData = async (id, checkIn, checkOut, guests, c) => {
    const url = buildRoomApiUrl(id, checkIn, checkOut, guests.adults, guests.children, guests.rooms, c);

    try {
        const response = await fetch(url);

        if (!response.ok) {
            return null;
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching hotel data:', error);
        return null;
    }
};

export default fetchRoomData;