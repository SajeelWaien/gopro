import {
  Table,
  TableCaption,
  Tbody,
  Td,
  Tfoot,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import React from "react";
import { useQuery } from "react-query";
import useListAgentsService from "../Services/ListAgents";

const ListAgents = () => {
  const { isLoading, isError, data, error } = useListAgentsService();
  console.log("data", data, isLoading, isError, error);
  if (isLoading) {
    return <div>Loading...</div>;
  } else if (isError) {
    return <div>Error: {error.message}</div>;
  }
  return (
    <Table variant='striped' colorScheme='brand'>
      <TableCaption>Valorant agents to date</TableCaption>
      <Thead>
        <Tr>
          <Th>Name</Th>
          <Th>Class</Th>
          <Th>Ult</Th>
        </Tr>
      </Thead>
      <Tbody>
        { data.map(item => (
          <Tr>
            <Td>{item.name}</Td>
            <Td>{item.class}</Td>
            <Td>{item.ult}</Td>
          </Tr>
        ))}
      </Tbody>
      <Tfoot>
        <Tr>
          <Th>Name</Th>
          <Th>Class</Th>
          <Th>Ult</Th>
        </Tr>
      </Tfoot>
    </Table>
  );
};

export default ListAgents;
