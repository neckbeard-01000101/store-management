import "./App.css";
import MainForm from "./routes/Form";
import { useEffect } from "react";
import { useLocation, Navigate } from "react-router-dom";
import Navbar from "./components/Navbar";
import Data from "./routes/Data";
import Profits from "./routes/Profits";
import Storage from "./routes/Storage";
import StorageForm from "./routes/StorageForm";

import { Route, Routes } from "react-router-dom";

function App() {
    function capitalize(str) {
        return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
    }
    const location = useLocation();
    useEffect(() => {
        const routeName = capitalize(
            location.pathname.replace("/", "") || "Form",
        );
        document.title = routeName;
    }, [location.pathname]);
    return (
        <>
            <Navbar />
            <Routes>
                <Route path="/form" element={<MainForm />} />
                <Route path="/data" element={<Data />} />
                <Route path="/profits" element={<Profits />} />
                <Route path="/storage" element={<Storage />} />
                <Route path="/add-to-storage" element={<StorageForm />} />
                <Route path="/" element={<Navigate to="/form" />} />
            </Routes>
        </>
    );
}

export default App;
