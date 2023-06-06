import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import Header from './componentes/Header/Header.jsx';
import Barra from './componentes/Barra/Barra.jsx';
import Login from './pages/login.jsx';
import Home from './pages/Home.jsx';
import Lista from './pages/Lista.jsx';
import Hotel from './pages/Hotel.jsx';
import SignUp from './pages/SignUp.jsx';

function App() {
  return (

    <Router>
      <Header />
        <Barra />
        <Routes>
        
          <Route path="/" element={<Home />}/>
          <Route path="/lista" element={<Lista />}/>
          <Route path="/hotel/:id" element={<Hotel />}/>
          <Route path="/login" element={<Login />}/>
          <Route path="/signup" element={<SignUp />}/>
        </Routes>
      </Router>

    
  );
}

export default App;
