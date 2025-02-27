import React, { useEffect, useState } from 'react';
import Popup from 'reactjs-popup';

const CreateStudent=()=>{
    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8080/api/students', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ id,}),
        });

        if (response.ok) {
            document.cookie = "token=your-authentication-token; path=/";
            onLogin(true);
        } else {
            
            alert('Unable to create new record');
            
        }
    };
    return(
        <div>
            <h1>Enter New Student Details</h1>
            <input type="text" placeholder='StudentName' />
            <input type="text" placeholder='Class' />
            <input type="text" placeholder='Address' />
            <button onClick={handleNewStud}>Save</button>

        </div>
    );

    
}
export default CreateStudent;