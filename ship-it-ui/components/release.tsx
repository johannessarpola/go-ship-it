import { Accordion, Badge, Button, Table } from "@mantine/core";
import ReleaseType from '../interfaces/release';

const Release = (props: ReleaseType) => {
    const { id, name, tag } = props
    console.log(props)
    return (<>
        <td>{name}</td>
        <td>{tag}</td>
    </>
    )
}

export default Release; 