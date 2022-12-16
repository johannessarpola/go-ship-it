import { Accordion, Badge, Table } from "@mantine/core";
import ReleaseType from '../interfaces/release';

const Release = ({ id, name, body, tag }: ReleaseType) => {
    return (
        <tr key={id}>
            <Accordion>
                <Accordion.Item value="row">
                    <Accordion.Control>
                        <td>{name}</td>
                        <td>{tag}</td>
                        <td>
                            <Badge>
                                Ship it 🚀
                            </Badge>
                        </td>
                        { body != null ? <Accordion.Panel>{body}</Accordion.Panel> : "" }
                    </Accordion.Control>
                </Accordion.Item>
            </Accordion>
        </tr>
    )
}

export default Release; 