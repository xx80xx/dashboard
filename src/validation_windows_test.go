// +build windows

package f2

import (
	"path/filepath"
	"testing"
)

func TestWindowsSpecificConflicts(t *testing.T) {
	testDir := setupFileSystem(t)

	table := []conflictTable{
		{
			name: "File name must not contain < or > characters",
			want: map[conflictType][]Conflict{
				invalidCharacters: {
					{
						source: []string{filepath.Join(testDir, "abc.pdf")},
						target: filepath.Join(testDir, "<>.pdf"),
						cause:  "<,>",
					},
				},
			},
			args: []string{"-f", "abc.pdf", "-r", "<>.pdf", testDir},
		},
		{
			name: "Directory or file name cannot contain trailing periods",
			want: map[conflictType][]Conflict{
				trailingPeriod: {
					{
						source: []string{
							filepath.Join(
								testDir,
								"No Pressure (2021) S1.E1.1080p.mkv",
							),
						},
						target: filepath.Join(
							testDir,
							`2021...\No Pressure (2021) S1.E1.1080p.mkv`,
						),
						cause: "",
					},
					{
						source: []string{
							filepath.Join(
								testDir,
								"No Pressure (2021) S1.E2.1080p.mkv",
							),
						},
						target: filepath.Join(
							testDir,
							`2021...\No Pressure (2021) S1.E2.1080p.mkv`,
						),
						cause: "",
					},
					{
						source: []string{
							filepath.Join(
								testDir,
								"No Pressure (2021) S1.E3.1080p.mkv",
							),
						},
						target: filepath.Join(
							testDir,
							`2021...\No Pressure (2021) S1.E3.1080p.mkv`,
						),
						cause: "",
					},
				},
			},
			args: []string{
				"-f",
				`.* \((2021)\) .*`,
				"-r",
				"$1.../{{f}}{{ext}}",
				testDir,
			},
		},
		{
			name: "File names must not contain :, |, or ? characters",
			want: map[conflictType][]Conflict{
				invalidCharacters: {
					{
						source: []string{filepath.Join(testDir, "abc.pdf")},
						target: filepath.Join(testDir, ":|?.pdf"),
						cause:  ":,|,?",
					},
				},
			},
			args: []string{"-f", "abc.pdf", "-r", ":|?.pdf", testDir},
		},
		{
			name: "File names must not be more than 260 characters",
			want: map[conflictType][]Conflict{
				maxLengthExceeded: {
					{
						cause:  "260 characters",
						source: []string{filepath.Join(testDir, "abc.pdf")},
						target: filepath.Join(
							testDir,
							"mSuhA166Od9QEPyZNr9YCGudIuxp7ousCVaTg4cNcOjuNjDZjKBNLJrjqwMhJhFFUM7Touvz054yah1hXkM7hKe6naBxg2FLfKO7YzdMpmgANN6yFF1jOyDGwPK7fn7uHymCG7NmpSXsS0QyURJMObjjTGfMKi6Zhd1l2fywyvyu0Oze8nNVdAve1HvXCyrqZJoczd5R84FIMDE2VK6zdM3D7Rbu6ZHj73tZnR476bc6q6pJiXiGaDzezx02Ngq6reI.c",
						),
					},
				},
			},
			args: []string{
				"-f",
				"abc.pdf",
				"-r",
				"mSuhA166Od9QEPyZNr9YCGudIuxp7ousCVaTg4cNcOjuNjDZjKBNLJrjqwMhJhFFUM7Touvz054yah1hXkM7hKe6naBxg2FLfKO7YzdMpmgANN6yFF1jOyDGwPK7fn7uHymCG7NmpSXsS0QyURJMObjjTGfMKi6Zhd1l2fywyvyu0Oze8nNVdAve1HvXCyrqZJoczd5R84FIMDE2VK6zdM3D7Rbu6ZHj73tZnR476bc6q6pJiXiGaDzezx02Ngq6reI.c",
				testDir,
			},
		},
	}

	runConflictCheck(t, table)
}

func TestWindowsNoConflict(t *testing.T) {
	testDir := setupFileSystem(t)

	cases := []testCase{
		{
			name: "Up to 260 emoji characters should not cause a conflict",
			want: []Change{
				{
					Source:  "abc.pdf",
					BaseDir: testDir,
					Target:  "😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀",
				},
			},
			args: []string{
				"-f",
				"abc.pdf",
				"-r",
				"😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀",
				testDir,
			},
		},
	}

	runFindReplace(t, cases)
}

func TestWindowsFixConflict(t *testing.T) {
	testDir := setupFileSystem(t)

	table := []testCase{
		{
			name: "Fix invalid characters present",
			want: []Change{
				{
					Source:  "abc.pdf",
					BaseDir: testDir,
					Target:  "name.pdf",
				},
			},
			args: []string{
				"-f",
				"abc.pdf",
				"-r",
				"name<>?|.pdf",
				"-F",
				testDir,
			},
		},
		{
			name: "Fix long file name conflict",
			want: []Change{
				{
					Source:  "abc.epub",
					BaseDir: testDir,
					Target:  "😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀",
				},
			},
			args: []string{
				"-f",
				"abc.epub",
				"-r",
				"😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀",
				"-F",
				testDir,
			},
		},
	}

	runFixConflict(t, table)
}
