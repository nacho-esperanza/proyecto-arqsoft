import logo from './logo.svg';
import './App.css';

import React, {useEffect, useState} from "react";

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Login from './pages/login';
import Home from './pages/Home';
import Lista from './pages/Lista';
import Hotel from './pages/Hotel';

function App() {
  return (

    <Router>
        <Routes>
          <Route path="/" element={<Home/>}></Route>
          <Route path="/lista" element={<Lista/>}></Route>
          <Route path="/lista/:id " element={<Hotel/>}></Route>
          <Route path="/login" element={<Login/>}></Route>
        </Routes>
      </Router>

    
  );
}

export default App;
