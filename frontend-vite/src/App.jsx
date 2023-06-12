import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import Header from './componentes/Header/Header.jsx';
import Barra from './componentes/Barra/Barra.jsx';
import Login from './pages/Login.jsx';
import Home from './pages/Home.jsx';
import Hotel from './pages/Hotel.jsx';
import SignUp from './pages/SignUp.jsx';
import Reserva from './pages/Reserva.jsx';
import MisReservas from './pages/MisReservas';
import Reservas from './pages/Reservas.jsx';
import Admin from './pages/Admin.jsx';

// Saque <Header /> para que no se vea en todas las paginas.

function App() {
  return (

    <Router>
        <Barra />
        <Routes>
        
          <Route path="/" element={<Home />}/>
          <Route path="/hotel/:id" element={<Hotel />}/>
          <Route path="/login" element={<Login />}/>
          <Route path="/signup" element={<SignUp />}/>
          <Route path="/reserva/:id" element={<Reserva />}/>
          <Route path="/misreservas" element={<MisReservas />}/>
          <Route path="/reservas" element={<Reservas />}/>
          <Route path="/admin" element={<Admin />}/>
        </Routes>
      </Router>

    
  );
}

export default App;
