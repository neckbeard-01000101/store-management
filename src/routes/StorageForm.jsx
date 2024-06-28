import axios from "axios";
import { HandleNumericInput } from "./Form";
async function handleSubmit(e) {
    e.preventDefault();
    if (!confirm("Are you sure you want to add a product?")) return;
    const formData = new FormData(e.target);
    const data = {
        type_of_product: formData.get("type-of-product"),
        product_color: formData.get("product-color"),
        product_size: formData.get("product-size"),
        quantity: parseInt(formData.get("quantity"), 10),
    };
    console.log("Sending data:", JSON.stringify(data));

    try {
        const response = await axios.post(
            "http://localhost:8000/add",
            JSON.stringify(data),
            {
                headers: {
                    "Content-Type": "application/json",
                },
            },
        );

        alert("product added successfully");
        console.log(response);
    } catch (err) {
        console.error("Error:", err.message);
    }
}
function StorageForm() {
    const fields = {
        select: [
            {
                labelName: "Clothes type",
                id: "type-of-product",
                options: [
                    { text: "Hoodie", value: "Hoodie" },
                    { text: "Slim T-shirt", value: "Slim_T-shirt" },
                    { text: "Hav T-shirt", value: "Hav_T-shirt" },
                ],
            },
            {
                labelName: "Color",
                id: "product-color",
                options: [
                    { text: "White", value: "White" },
                    { text: "Black", value: "Black" },
                ],
            },
            {
                labelName: "Size",
                id: "product-size",
                options: [
                    { text: "S", value: "S" },
                    { text: "M", value: "M" },
                    { text: "L", value: "L" },
                    { text: "XL", value: "XL" },
                    { text: "XXL", value: "XXL" },
                    { text: "XXXL", value: "XXXL" },
                    { text: "XXXXL", value: "XXXXL" },
                ],
            },
        ],
        inputs: [
            {
                labelName: "Quantity",
                id: "quantity",
                isNumeric: true,
                placeHolder: "Quantity",
                inputType: "text",
            },
        ],
    };

    return (
        <form className="storage-form" onSubmit={(e) => handleSubmit(e)}>
            {fields.select.map((item) => (
                <div className="form-element" key={item.labelName}>
                    <label htmlFor={item.id}>{item.labelName}</label>
                    <select name={item.id} id={item.id}>
                        {item.options.map((option) => (
                            <option key={option.value} value={option.value}>
                                {option.text}
                            </option>
                        ))}
                    </select>
                </div>
            ))}
            {fields.inputs.map((item) => (
                <div key={item.id} className="form-element">
                    <label htmlFor={item.id}>{item.labelName}</label>
                    <input
                        type={item.inputType}
                        id={item.id}
                        name={item.id}
                        onChange={
                            item.isNumeric ? (e) => HandleNumericInput(e) : null
                        }
                    />
                </div>
            ))}
            <button type="submit">Add</button>
        </form>
    );
}

export default StorageForm;
