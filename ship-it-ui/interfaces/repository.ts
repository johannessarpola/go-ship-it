import type ReleaseType from './release'

type RepositoryType = {
    owner: string
    name: string,
    releases: ReleaseType[],
}

export default RepositoryType;