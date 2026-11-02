---
title: "Chronicles Privacy Policy"
date: 2026-11-01T00:00:00-08:00
# Unlisted: App Store Connect links here. Not in nav, the post list, RSS, or the sitemap.
_build:
  list: never
---

Effective November 1, 2026. This policy covers Chronicles for iPhone and Mac, made by Christopher Loverich ([pinecoder.dev](https://pinecoder.dev)).

## What Chronicles collects

Nothing. There are no accounts, analytics, ads, tracking, or third-party SDKs. I don't run a server for Chronicles, and your notes never reach me.

## On your device

Your notes, attachments, and settings are stored on your device, in the app's own storage. They stay there unless you turn on one of the features below.

## iCloud sync

Sync is optional and off until you turn it on. It uses your own private iCloud database (Apple's CloudKit), under your Apple Account. I have no access to it.

A new library is encrypted on your device before anything goes to iCloud: notes, titles, journal names, tags, and attachments. Only your devices can read it. The key is kept in your iCloud Keychain, so your other devices signed in to the same Apple Account can open the library. When you set up sync, Chronicles also shows you a recovery code, which only you hold.

To delete the copy in iCloud, use Settings › iCloud › Remove Library from iCloud. This deletes the library from iCloud and retires its key; the library on each of your devices stays as it is.

## Camera and microphone

Chronicles uses the camera only when you take a photo or video to add to a note, and the microphone only to record sound with a video you're taking. What you capture is saved in your library like any other attachment.

## Export, backups, and blog publishing

Export (Settings › Import & export) writes your notes where you choose: as plain Markdown (a folder or a `.chronicles.zip` file), or as an encrypted file that opens with your library's key or recovery code. Backups, if you set them up, are written to a folder you pick. Blog publishing, an experimental Labs feature that is off by default, writes the notes you publish to a folder you pick. What happens to those files afterward depends on where you put them.

## Assistants (Mac only)

Assistant access is off by default. If you turn it on and connect an assistant such as Claude, that assistant can read the journals you check, and, if you allow it, add or edit notes in them. If you also turn on Share diagnostics, it can read the app's version, settings, and sync log.

The assistant app starts Chronicles on your Mac and talks to it directly; Chronicles makes no network connections for this. Once an assistant has read your notes, what happens to them is governed by that assistant provider's privacy policy, not this one.

## Crash reports

Chronicles has no crash reporting of its own. Apple shares crash reports with me only if you turn on sharing with app developers in your device's Analytics & Improvements settings.

## Children

Chronicles is not directed at children under 13, and it collects no information from anyone.

## Source code

Chronicles is open source, under the GNU Affero General Public License, version 3 or later.
<!-- Source link goes here once the repo is public: https://github.com/cloverich/chronicles -->

## Changes

If this policy changes, I'll update this page and its effective date.

## Contact

Questions? Email [support@pinecoder.dev](mailto:support@pinecoder.dev).
