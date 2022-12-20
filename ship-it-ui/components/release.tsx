import { Accordion, Badge, Button, Table } from "@mantine/core";
import ReleaseType from '../interfaces/release';

const Release = ({ id, name, tag }: ReleaseType) => {
    return (<>
        <td>{name}</td>
        <td>{tag}</td>
    </>
    )
}

export default Release; 