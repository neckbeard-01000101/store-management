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
    async function handleQuantityChange(e, index) {
        const inputField = document.getElementById(`quantity-updater-${index}`);
        const parameters = {
            oper: e.target.innerText === "+" ? "add" : "sub",
            amount: inputField.value,
            type: storage[index]["type_of_product"],
            color: storage[index]["product_color"],
            size: storage[index]["product_size"],
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

    useEffect(() => {
        fetchData();
    }, []);
    return (
        <div className="profits">
            <h2>Storage</h2>
            <table>
                <thead>
                    <tr>
                        <th>Product</th>
                        <th>Color</th>
                        <th>Size</th>
                        <th>Quantity</th>
                        <th>Update quantity</th>
                    </tr>
                </thead>
                <tbody>
                    {storage?.map((item, index) => (
                        <tr key={index}>
                            <td>{item["type_of_product"]}</td>
                            <td>{item["product_color"]}</td>
                            <td>{item["product_size"]}</td>
                            <td>{item.quantity}</td>
                            <td>
                                <input
                                    type="text"
                                    id={`quantity-updater-${index}`}
                                    placeholder="Enter quantity to add or sub"
                                    onChange={(e) => HandleNumericInput(e)}
                                />
                                <br />
                                <div className="btn-container">
                                    <button
                                        onClick={(e) =>
                                            handleQuantityChange(e, index)
                                        }
                                    >
                                        +
                                    </button>
                                    <button
                                        onClick={(e) =>
                                            handleQuantityChange(e, index)
                                        }
                                    >
                                        -
                                    </button>
                                </div>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
export default Storage;
