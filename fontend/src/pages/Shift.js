import {
  Link,
  Switch,
  Route,
  Redirect,
  useParams,
  useRouteMatch,
  useHistory
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
          <FormShift type="edit"/>
        </Route>
    </Switch>
  </>
}


function FormShift(props){
  let history = useHistory()
  let {user, type} = props
  let { id } = useParams();
  let [start_date, setStartDate] = useState("");
  let [end_date, setEndDate] = useState("");
  let [users, setUsers] = useState([]);
  let [user_id, setUserId] = useState("");

  let getUsers = () => {
    axios.get('/user').then(res => {
      setUsers(res.data.data)
    }).catch(err => {
      NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
    })
  }

  let getShift = () => {
    axios.get('/shift', {
      params: {
        id
      }
    }).then(res => {
      if(res.data.data.length){
        console.log("update cuy", res.data.data)
        setStartDate(res.data.data[0].start_date)
        setEndDate(res.data.data[0].end_date)
        setUserId(res.data.data[0].assign_user_id)
      } else {
        history.push("/shift")
        NotificationManager.error("not found", '', 3000);
      }
      
    }).catch(err => {
      NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
    })
  }

  useEffect(() => {
    getUsers()
    if(type == 'edit'){
      getShift()
    }
  }, [])

  let submitShift = () => {
    let data = {
      start_date : start_date ? new Date(
        start_date.substr(0,4),
        start_date.substr(5,2),
        start_date.substr(8,2),
        start_date.substr(11,2),
        start_date.substr(14,2),
      ).toISOString() : "",
      end_date : end_date ? new Date(
        end_date.substr(0,4),
        end_date.substr(5,2),
        end_date.substr(8,2),
        end_date.substr(11,2),
        end_date.substr(14,2),
      ).toISOString(): "",
      assign_user_id: user_id ? user_id :  user.id
    }

   if(type == "edit") {
    axios.put('/shift/' + id , data).then(res => {
      NotificationManager.success('update success', '', 3000);
    }).catch(err => {
      NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
    })
   } else {
      axios.post('/shift', data).then(res => {
        NotificationManager.success('add success', '', 3000);
      }).catch(err => {
        NotificationManager.error(err.response ? err.response.data : 'error : something not right', '', 3000);
      })
   }
  }

  return <div className="form-shift">
    <div style={{padding: "8px", marginBottom: "0px"}}>
      Assign : 
      <select onChange={e => {setUserId(e.target.value)}} className="select-user border-input">
        {users.map((v,i) => {
          return <option value={v.id}>{v.name}</option>
        })}
      </select>
    </div>
    <div>Start Date : <input value={start_date ? new Date(new Date(start_date).toLocaleString("en-US", {timeZone: 'UTC'})).toISOString().substr(0,16) : null} className="border-input" onChange={e => setStartDate(e.target.value)} type="datetime-local"/></div>
    <div>End Date : <input value={end_date ? new Date(new Date(end_date).toLocaleString("en-US", {timeZone: 'UTC'})).toISOString().substr(0,16) : null} className="border-input" onChange={e => setEndDate(e.target.value)} type="datetime-local"/></div>
    <button className="button-primary" onClick={() => submitShift()}>Add New Shift</button>
  </div>
}

function ListShift(props){
  let history = useHistory()
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

  let publishShift = (id, data) => {
    // eslint-disable-next-line
   if(confirm("publish ? ")) {
    axios.put('/shift/'+id, {
      start_date : data.start_date,
      end_date : data.end_date,
      assign_user_id : data.assign_user_id,
      status : 'published'
    })
    .then(res => {
      NotificationManager.success('publish success', '', 3000);
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
        <h4>{props.user.name} ({v.status})</h4>
        <div style={{display:"flex", justifyContent: "space-between"}}>
          <StartEnd start_date={v.start_date} end_date={v.end_date}/>
          <div>
            <button onClick={() => history.push("/shift/" +v.id)}>Edit</button>
            <button onClick={() => deleteShift(v.id)}>Delete</button>
            <button onClick={() => publishShift(v.id, v)}>Publish</button>
          </div>
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