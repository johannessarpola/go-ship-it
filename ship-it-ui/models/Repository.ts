import Release from "../models/Release";

type Repository = {
    name: string
    releases: Release[]
}

export default Repository;