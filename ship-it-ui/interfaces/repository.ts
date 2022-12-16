import type ReleaseType from './release'

type RepositoryType = {
    name: string,
    releases: ReleaseType[],
}

export default RepositoryType;