import React from "react";

import '../App.css';

import {useParams} from "react-router-dom";
  
const Home = () => {
  return (
    <div>
      <a>
      Home
      </a>
      <a href="/login">
        login
      </a>
    </div>
    
  );
};
  
export default Home;