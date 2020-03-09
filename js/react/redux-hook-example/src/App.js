import React,{useState} from 'react';
import {useDispatch,useSelector} from 'react-redux';
import {addCount,subCount} from './store/actions/index';

//uesState
// function App(){
//     const [count,setCount]=useState(0);
//     return (
//       <div>
//         <p>Count:{count}</p>
//         <button onClick={()=>setCount(count+1)}>Add</button>
//       </div>
//     )
// }

function App() {
  const count=useSelector(state=>state.data);
  const dispatch=useDispatch()

  return(
    <main>
      <div>Count :{count}</div>
      <button onClick={()=>dispatch(addCount())}>Add to count</button>
      <button onClick={()=>dispatch(subCount())}>Sub to count</button>
    </main>
  )
}

export default App;
