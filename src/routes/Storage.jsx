import axios from "axios";
import { HandleNumericInput } from "./Form";
import { useEffect, useState } from "react";
function Storage() {
    const [storage, setStorage] = useState([]);

    async function fetchData() {
        try {
            const response = await axios.get(`${SERVER_URL}/get/storage`);
            setStorage(response.data);
        } catch (err) {
            console.log(err);
        }
    }
    const SERVER_URL = "http://localhost:8000";
    async function handleQuantityChange(e, index, item, title) {
        const inputFieldId = `quantity-updater-${title}-${index}`;
        const inputField = document.getElementById(inputFieldId);
        const parameters = {
            oper: e.target.innerText === "+" ? "add" : "sub",
            amount: inputField.value,
            type: item["type_of_product"],
            color: item["product_color"],
            size: item["product_size"],
        };
        console.log(
            `${SERVER_URL}/updateAmount/?type=${parameters.type}&amount=${parameters.amount}&oper=${parameters.oper}&color=${parameters.color}&size=${parameters.size}`,
        );
        try {
            if (!confirm("Are you sure you want to update the quantity?"))
                return;
            await axios.post(
                `${SERVER_URL}/updateAmount/?type=${parameters.type}&amount=${parameters.amount}&oper=${parameters.oper}&color=${parameters.color}&size=${parameters.size}`,
            );
            fetchData();
        } catch (err) {
            console.log(err);
        }
    }
    async function deleteItem(itemId) {
        try {
            if (!confirm("Are you sure you want to delete this item?")) return;
            await axios.delete(
                `${SERVER_URL}/deleteDocument/storage/${itemId}`,
            );
            fetchData();
            alert("Item deleted successfully");
        } catch (err) {
            console.error("Failed to delete item:", err);
            alert(err.response.data);
        }
    }
    useEffect(() => {
        fetchData();
    }, []);
    function sortFunction(a, b) {
        const sizeOrder = ["S", "M", "L", "XL", "XXL", "XXXL", "XXXXL"];

        if (a.type_of_product < b.type_of_product) return -1;
        if (a.type_of_product > b.type_of_product) return 1;

        if (a.product_color > b.product_color) return -1;
        if (a.product_color < b.product_color) return 1;

        const sizeIndexA = sizeOrder.indexOf(a.product_size);
        const sizeIndexB = sizeOrder.indexOf(b.product_size);
        if (sizeIndexA < sizeIndexB) return -1;
        if (sizeIndexA > sizeIndexB) return 1;

        return a.quantity - b.quantity;
    }
    const hoodies = storage
        ?.filter((item) => item.type_of_product === "Hoodie")
        .sort(sortFunction);
    const havTshirts = storage
        ?.filter((item) => item.type_of_product === "Hav_T-shirt")
        .sort(sortFunction);
    const slimTshirts = storage
        ?.filter((item) => item.type_of_product === "Slim_T-shirt")
        .sort(sortFunction);
    function renderTable(items, title) {
        return (
            <div className="product-table">
                <h2>{title}</h2>
                <table>
                    <thead>
                        <tr>
                            <th>Product</th>
                            <th>Color</th>
                            <th>Size</th>
                            <th>Quantity</th>
                            <th>Update quantity</th>
                            <th>Delete</th>
                        </tr>
                    </thead>
                    <tbody>
                        {items.map((item, index) => (
                            <tr key={index}>
                                <td>{item["type_of_product"]}</td>
                                <td>{item["product_color"]}</td>
                                <td>{item["product_size"]}</td>
                                <td>{item.quantity}</td>
                                <td>
                                    <input
                                        type="text"
                                        id={`quantity-updater-${title}-${index}`}
                                        placeholder="Enter quantity to add or sub"
                                        onChange={(e) => HandleNumericInput(e)}
                                    />
                                    <br />
                                    <div className="btn-container">
                                        <button
                                            onClick={(e) =>
                                                handleQuantityChange(
                                                    e,
                                                    index,
                                                    item,
                                                    title,
                                                )
                                            }
                                        >
                                            +
                                        </button>
                                        <button
                                            onClick={(e) =>
                                                handleQuantityChange(
                                                    e,
                                                    index,
                                                    item,
                                                    title,
                                                )
                                            }
                                        >
                                            -
                                        </button>
                                    </div>
                                </td>
                                <td>
                                    <button
                                        className="delete-btn"
                                        onClick={() => deleteItem(item._id)}
                                    >
                                        Delete
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        );
    }

    return (
        <div className="storage">
            {renderTable(hoodies, "Hoodies")}
            {renderTable(havTshirts, "Hav T-Shirts")}
            {renderTable(slimTshirts, "Slim T-Shirts")}
        </div>
    );
}

export default Storage;
