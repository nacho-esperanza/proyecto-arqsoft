import { lazy } from "react";
import Home from "pages/Home.jsx"
import Lista from "pages/Lista.jsx";
import Login from "pages/Login.jsx";
import SignUp from "pages/SignUp.jsx";
import Hotel from "pages/Hotel.jsx";
import Reserva from "pages/Reserva.jsx";
import MisReservas from "pages/MisReservas.jsx";
import Reservas from "pages/Reservas.jsx";
import Admin from "pages/Admin.jsx";
import CrearHotel from "pages/CrearHotel.jsx";
import AñadirImagenes from "pages/AñadirImagenes.jsx";


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
        path: `/signup`,
        Element: SignUp,
    },
    {
        id: 3,
        path: `/hotel/:id`,
        Element: Hotel,
    },
    {
        id: 4,
        path: `/reserva/:id`,
        Element: Reserva,
    },
    {
        id: 5,
        path: `/misreservas`,
        Element: MisReservas,
    },
    {
        id: 6,
        path: `/reservas`,
        Element: Reservas,
    },
    {
        id: 7,
        path: `/admin`,
        Element: Admin,
    },
    {
        id: 8,
        path: `/crearhotel`,
        Element: CrearHotel,
    },
    {
        id: 9,
        path: `/añadirimagenes`,
        Element: AñadirImagenes,
    },



];