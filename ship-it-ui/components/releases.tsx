import Release from "./release";
import ReleaseType from "../interfaces/release"

interface Props {
    releases: ReleaseType[]
}

const Releases = ({ releases }: Props) => {
    return (releases.map((release, idx) => {
            return (<Release key={idx} {...release}></Release>)
        })
    )
}

export default Releases;