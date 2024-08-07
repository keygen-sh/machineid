// +build linux

package machineid

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"

	// on Docker, there are no files for machine-id, and installing
	// dbus creates a static machine-id for all containers.
	// To overcome this problem, we can add a last fallback value
	// which is the hostname file, which, in Docker is the container
	// name.
	hostnamePath = "/etc/hostname"
)

// machineID returns the uuid specified at `/var/lib/dbus/machine-id` or `/etc/machine-id`.
// If there is an error reading the files an empty string is returned.
// See https://unix.stackexchange.com/questions/144812/generate-consistent-machine-unique-id
func machineID() (string, error) {
	id, err := readFile(dbusPath)
	if err != nil {
		// try fallback path
		id, err = readFile(dbusPathEtc)
	}
	if err != nil {
		// this might be a docker container, use the hostname instead
		id, err = readFile(hostnamePath)
		if err != nil {
			// we tried all fallbacks
			return "", err
		}
	}
	return trim(string(id)), nil
}
