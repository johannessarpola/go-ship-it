import Release from "./release";
import ReleaseType from "../interfaces/release"
import { Table } from "@mantine/core";

interface Props {
    releases: ReleaseType[]
}

const Releases = ({ releases }: Props) => {
    return (
        <Table>
            <thead>
                <tr>
                    <td>Name</td>
                    <td>Tag</td>
                </tr>
            </thead>
            <tbody>
                {releases.map((release, idx) => {
                    return (<Release key={idx} {...release}></Release>)
                })}
            </tbody>

        </Table>

    )
}

export default Releases;