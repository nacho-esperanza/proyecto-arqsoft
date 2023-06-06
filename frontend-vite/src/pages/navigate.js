import { lazy } from "react";
import Home from "pages/Home.jsx"
import Lista from "pages/Lista.jsx";
import Login from "pages/Login.jsx";
import SingUp from "pages/SignUp.jsx";
import Hotel from "pages/Hotel.jsx";


const Home = lazy(() => import("pages/Home.jsx"));
//estas son las rutas de las paginas de nuestro programa

export const navigation = [
    {
        id: 0,
        path: "/",
        Element: Home,
    },
    {
        id: 1,
        path: '/login',
        Element: Login,
    },
    {
        id: 2,
        path: `/singup`,
        Element: SingUp,
    },
    {
        id: 3,
        path: '/lista',
        Element: Lista,
    },
    {
        id: 4,
        path: `/hotel/'${id}`,
        Element: Hotel,
    },
];