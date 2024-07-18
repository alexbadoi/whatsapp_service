const buildHotelApiUrl = (id, checkIn, checkOut, adult, children, room, filters) => {
    const baseUrl = 'http://localhost:8080/api/hotels';
    const params = new URLSearchParams({
        id,
        checkIn,
        checkOut,
        adult,
        room
    });
    if (children.length > 0) {
        if (children.length === 1) {
            params.append('children', children[0]);
        } else {
            params.append('children', children.join(','));
        }
    }
    if (filters.starRating.length > 0) {
        params.append('starRating', filters.starRating.join(','));
    }
    if (filters.guestRating.length > 0) {
        params.append('guestRating', filters.guestRating.join(','));
    }
    if (filters.propertyFacilities.length > 0) {
        params.append('facilities', filters.propertyFacilities.join(','));
    }
    if (filters.bedrooms.length > 0) {
        params.append('bedrooms', filters.bedrooms.join(','));
    }
    if (filters.paymentOptions.length > 0) {
        params.append('paymentOptions', filters.paymentOptions.join(','));
    }
    if (filters.popularWithFamilies) {
        params.append('kidsStayForFree', '250691');
    }
    return `${baseUrl}?${params.toString()}`;
};

const fetchHotelData = async (id, checkIn, checkOut, adult, children, room, filters) => {
    const url = buildHotelApiUrl(id, checkIn, checkOut, adult, children, room, filters);

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

export default fetchHotelData;