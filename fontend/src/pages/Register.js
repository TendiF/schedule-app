import {
  Link,
} from "react-router-dom";
export default function Login() {
  return <>
    <div className="header">
      <h4>Register</h4>
    </div>
    <div>
      <div className="login">
        <div>
          <p>Name</p>
          <div className="input" style={{alignItems:'center'}} >
            <input placeholder="Name" style={{height:'30px', marginLeft:'10px'}}/>
          </div>
        </div>
      <button style={{marginTop: '15px'}} className="button-primary">Register</button>
      <Link style={{marginTop:'10px'}} to="login">Already have account ? <span style={{color:'blue'}}>go to login</span> </Link>
      </div>
    </div>
  </>
}