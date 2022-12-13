type Props = {
    repos: Repository[]
}

type Repository = {
    name: string
    releases: Release[]
}

type Release = {
    id: string,
    name: string,
    body: string,
    tag: string
}

const RepoList = (props: Props) => {
    const abc = 123;
    return (<table className="table-auto">
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Body</th>
                <th>tag</th>
            </tr>
        </thead>
        <tbody>
            rep
        </tbody>
    </table>)
}

export default RepoList;