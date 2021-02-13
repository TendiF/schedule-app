import {
  Link,
} from "react-router-dom"

import { useState } from 'react'
import { useDispatch } from 'react-redux'
import { doRegister } from '../app/userReducer'

export default function Login() {
  const dispatch = useDispatch();
  const [name, setName] = useState("");

  return <>
    <div className="header">
      <h4>Register</h4>
    </div>
    <div>
      <div className="login">
        <div>
          <p>Name</p>
          <div className="input" style={{alignItems:'center'}} >
            <input value={name} onChange={e => setName(e.target.value)} placeholder="Name" style={{height:'30px', marginLeft:'10px'}}/>
          </div>
        </div>
      <button onClick={() => dispatch(doRegister({name}))} style={{marginTop: '15px'}} className="button-primary">Login</button>
      <Link style={{marginTop:'10px'}} to="login">Already have account ? <span style={{color:'blue'}}>go to login</span> </Link>
      </div>
    </div>
  </>
}