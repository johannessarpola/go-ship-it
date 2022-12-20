import Repository from "./repository";
import RepositoryType from "../interfaces/repository"

interface Props {
    repos: RepositoryType[],
}


const Repositories = ({ repos }: Props) => {
    return (
        <>
            {
                repos.map((repo: RepositoryType, idx: any ) => {
                    return (
                        <Repository key={idx} owner={repo.owner} name={repo.name} releases={repo.releases}></Repository>
                    );
                })
            }
        </>
    )
}

export default Repositories;