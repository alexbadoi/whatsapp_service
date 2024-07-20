const SaveProposal= async (data) => {
    try {
        var res = await fetch('/api/proposal', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });
        if (res.ok) {
            return 1;
        }else{
            return null;
        }
    } catch (error) {
        console.log(error);
        return null;
    }
};

export default SaveProposal;