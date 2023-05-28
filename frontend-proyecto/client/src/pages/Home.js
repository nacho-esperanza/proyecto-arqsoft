import React from "react";

import '../App.css';

import {useParams} from "react-router-dom";
  
const Home = () => {
  return (
    <div>
      Home
      <a href="/Login">
        login
      </a>
    </div>
    
  );
};
  
export default Home;