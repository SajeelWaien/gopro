import React from 'react';
import { useQuery } from 'react-query';
import useListAgentsService from '../Services/ListAgents';

const ListAgents = () => {
  const  { isLoading, isError, data, error } = useListAgentsService()
console.log("data", data, isError, error);
  if(isLoading) {
    return <div>Loading...</div>
  }
  else if(isError) {
    return <div>Error: {error.message}</div>
  }
  return (
    <div>
      <ul>
        { data.map(item => <li>{item.name}</li>) }
      </ul>
    </div>
  )
};

export default ListAgents;
