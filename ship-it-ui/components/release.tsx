import { Badge, Table } from "@mantine/core";
import ReleaseType from '../interfaces/release';


const Release = ({ id, name, body, tag }: ReleaseType) => {
    const abc = 123;
    return (
        <tr key={id}>
            <td>{name}</td>
            <td>{tag}</td>
            <td>
                <Badge>
                    Ship it 🚀
                </Badge>
            </td>

        </tr>

    )
}

export default Release; 