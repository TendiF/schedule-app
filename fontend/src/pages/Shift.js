import {
  Link,
  Switch,
  Route,
  Redirect,
  useParams,
  useRouteMatch
} from "react-router-dom"
import {NotificationManager} from 'react-notifications';
import { useSelector } from 'react-redux';
import { useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'

import axios from '../app/utils/axios'
import { getUser } from '../app/userReducer'
import { getAxiosShift, getShift } from '../app/shiftReducer'

export default function Shift() {
  const {user} = useSelector(getUser);
  let { path, url } = useRouteMatch();
  return <>
    {!user.data.id && <Redirect to="/login" />} 
    <div className="header">
      <h4>Shift {user.data ? user.data.id : ""}</h4>
    </div>
    <Switch>
        <Route exact path={path}>
         <ul>
            <li>
              <Link to={`${url}/form`}>Add Shift</Link>
            </li>
            <li>
              <Link to={`${url}/1`}>Detail </Link>
            </li>
          </ul>
          <ListShift user={user.data}/>
        </Route>
        <Route path={`${path}/form`}>
          <FormShift user={user.data} />
        </Route>
        <Route path={`${path}/:id`}>
          <DetailShift />
        </Route>
    </Switch>
  </>
}

function DetailShift(){
  let { id } = useParams();
  return <div>
    Detail Shift
  </div>
}

function FormShift(props){
  let {user} = props
  let { id } = useParams();
  let [start_date, setStartDate] = useState("");
  let [end_date, setEndDate] = useState("");

  let submitShift = () => {
    let data = {
      start_date : new Date(start_date).toISOString(),
      end_date : new Date(end_date).toISOString(),
      assign_user_id: user.id
    }
    axios.post('/shift', data).then(res => {
      NotificationManager.success('add success', '', 3000);
    }).catch(err => {
      NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
    })
  }

  return <div className="form-shift">
    <div>Issued : 
      <select>
        <option>Siapa</option>
      </select>
    </div>
    <div>Start Date : <input onChange={e => setStartDate(e.target.value)} type="datetime-local"/></div>
    <div>End Date : <input onChange={e => setEndDate(e.target.value)} type="datetime-local"/></div>
    <button onClick={() => submitShift()}>Add New Shift</button>
  </div>
}

function ListShift(props){
  const dispatch = useDispatch();
  const {shift} = useSelector(getShift);

  useEffect(() => {
    dispatch(getAxiosShift({id_user : props.user.id}))
  }, [dispatch])

  let deleteShift = id => {
    // eslint-disable-next-line
   if(confirm("delete ? ")) {
    axios.delete('/shift/'+id)
    .then(res => {
      NotificationManager.success('delete success', '', 3000);
      dispatch(getAxiosShift({id_user : props.user.id}))
    })
    .catch(err => {
      NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
    })
   }
  }

  return <div>
    {Array.isArray(shift.data) && !shift.data.length && "empty data"}
    {Array.isArray(shift.data) && shift.data.map((v,i) => {
      return <div key={v.start_date} className="shiftCard">
        <h4>Name</h4>
        <div style={{display:"flex", justifyContent: "space-between"}}>
          <StartEnd start_date={v.start_date} end_date={v.end_date}/>
          <button onClick={() => deleteShift(v.id)}>Delete</button>
        </div>
      </div>
    })}
  </div>
}

let StartEnd = (props) => {
  let {start_date, end_date} = props
  let comp = null
  if(new Intl.DateTimeFormat('en-GB', { dateStyle: 'short'}).format(new Date(end_date)) === new Intl.DateTimeFormat('en-GB', { dateStyle: 'short'}).format(new Date(start_date))){
    comp = <p>
      {new Intl.DateTimeFormat('en-GB', { dateStyle: 'short', timeStyle: 'short' }).format(new Date(start_date))} - {new Intl.DateTimeFormat('en-GB', { timeStyle: 'short' }).format(new Date(end_date))}
      </p>
  }else {
    comp = <p>
    {new Intl.DateTimeFormat('en-GB', { dateStyle: 'short', timeStyle: 'short' }).format(new Date(start_date))} - {new Intl.DateTimeFormat('en-GB', { dateStyle: 'short', timeStyle: 'short' }).format(new Date(end_date))}
    </p>
  }

  return <>{comp}</>
}