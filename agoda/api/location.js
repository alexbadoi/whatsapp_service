const fetchLocations = async (inputValue) => {
    try {
        const response = await fetch(`/api/locations?name=${inputValue}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return await response.json();
    } catch (error) {
        console.error('Error fetching locations:', error);
        return { data: [] };
    }
};

const parseLocations = (data) => {
    return data.data.map(item => ({
        value: item.id,
        label: item.name,
    }));
};

export {fetchLocations, parseLocations};