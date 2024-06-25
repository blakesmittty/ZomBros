import React, { useEffect, useState } from 'react';

function Items({ playerId }) {
    const [items, setItems] = useState([]);

    useEffect(() => {
        const fetchItems = async () => {
            const response = await fetch(`/getUnlockedItems?player_id=${playerId}`);
            const data = await response.json();
            setItems(data);
        };
        fetchItems();
    }, [playerId]);

    return (
        <div>
            <h2>Unlocked Items</h2>
            <ul>
                {items.map((item) => {
                    <li key={item.id}>{item.name}</li>
                })}
            </ul>
        </div>
    );
}

export default Items;