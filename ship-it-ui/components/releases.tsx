import Release from "./release";
import ReleaseType from "../interfaces/release"
import { Button, Table } from "@mantine/core";

interface Props {
    repoName: string
    repoOwner: string
    releases: ReleaseType[]
}

async function rocket(repoName: string, repoOwner: string, tag: string) {
    const rocket = {
        name: repoName,
        owner: repoOwner,
        tag: tag
    }
    fetch('/api/ship-it', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(rocket),
    })
}


const Releases = ({ repoOwner, repoName, releases }: Props) => {
    return (
        <Table>
            <thead>
                <tr>
                    <td>Name</td>
                    <td>Tag</td>
                    <td></td>
                </tr>
            </thead>
            <tbody>

                {releases.map((release, idx) => {
                    return (
                        <tr key={idx}>
                            <Release key={idx} {...release}></Release>
                            <td>
                                <Button onClick={() => rocket(repoName, repoOwner, release.tag)} variant="gradient" gradient={{ from: '#ed6ea0', to: '#ec8c69', deg: 35 }}>
                                    Ship it 🚀
                                </Button>
                            </td>
                        </tr>
                    )
                })}
            </tbody>

        </Table>

    )
}

export default Releases;