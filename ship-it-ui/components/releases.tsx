import Release from "./release";
import ReleaseType from "../interfaces/release"

type Props = {
    releases: ReleaseType[]
}

const Releases = ({ releases }: Props) => {
    return (
        releases.map((release, idx) => {
            return <Release {...release}></Release>
        })
    )
}

export default Releases;