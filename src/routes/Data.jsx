import React, { useEffect, useState } from 'react';

const SERVER_URL = "http://localhost:8000";
export default function Data() {
    const [collections, setCollections] = useState([]);
    const [data, setData] = useState([]);
    const [selectedCollection, setSelectedCollection] = useState('06/2024');

    useEffect(() => {
        const fetchCollections = async () => {
            const response = await fetch(`${SERVER_URL}/collections`);
            const data = await response.json();
            setCollections(data);
        };

        fetchCollections();
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch(`${SERVER_URL}/get/${selectedCollection}`);
            const data = await response.json();
            setData(data);
        };

        if (selectedCollection) {
            fetchData();
        }
    }, [selectedCollection]);

    const handleChange = (event) => {
        setSelectedCollection(event.target.value);
    };

    return (
        <div className='data-page'>
            <label htmlFor="month">Select a month: </label>
            <select id="month" value={selectedCollection} onChange={handleChange}>
                {collections.map(collection => (
                    <option key={collection} value={collection}>
                        {collection}
                    </option>
                ))}
            </select>
            {selectedCollection && (
                <div>
                    <h2>{selectedCollection}</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>Order Number</th>
                                <th>Order State</th>
                                <th>Customer Name</th>
                                <th>Customer City</th>
                                <th>Customer Phone</th>
                                <th>Seller Name</th>
                                <th>Total Cost</th>
                                <th>Seller Profit</th>
                                <th>Delivery Fee</th>
                                <th>Cost of Product</th>
                                <th>Size</th>
                                <th>Color</th>
                                <th>Clothes Type</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                            data.map((item, index) => (
                                <tr key={index}>
                                    <td>{item['order-num']}</td>
                                    <td>{item['order-state']}</td>
                                    <td>{item['customer-name']}</td>
                                    <td>{item['customer-city']}</td>
                                    <td>{item['customer-phone']}</td>
                                    <td>{item['seller-name']}</td>
                                    <td>{item['total-cost']}</td>
                                    <td>{item['seller-profit']}</td>
                                    <td>{item['delivery-fee']}</td>
                                    <td>{item['cost-of-product']}</td>
                                    <td>{item.size}</td>
                                    <td>{item.color}</td>
                                    <td>{item['clothes-type']}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
}